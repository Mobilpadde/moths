package token

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math"
	"time"

	"github.com/Mobilpadde/moths/v6/token/code"
)

func (g *Generator) Next() (code.Code, error) {
	token, err := g.getToken()
	if err != nil {
		return code.Code{}, err
	}

	truncated := uint64(g.timing.time.Unix() / int64(g.period.Seconds()))
	period := time.Unix(int64(truncated)*int64(g.period.Seconds()), 0)

	return code.NewCode(token, g.amount, g.emojies, period, period.Add(g.period))
}

func (g *Generator) getToken() (string, error) {
	g.timing.curr = time.Now().UTC()

	since := g.timing.curr.Sub(g.timing.last)
	if since < g.period {
		return g.lastToken, nil
	}

	g.timing.last = g.timing.curr
	g.timing.time = g.timing.time.Add(since)
	period := uint64(g.timing.time.Unix() / int64(g.period.Seconds()))

	// https://github.com/pquerna/code/blob/master/hotp/hotp.go#L95-L123
	buf := make([]byte, 8)
	h := hmac.New(sha1.New, g.secret)
	binary.BigEndian.PutUint64(buf, period)

	h.Write(buf)
	sum := h.Sum(nil)

	// "Dynamic truncation" in RFC 4226
	// http://tools.ietf.org/html/rfc4226#section-5.4
	offset := sum[len(sum)-1] & 0xf
	value := int64(((int(sum[offset]) & 0x7f) << 24) |
		((int(sum[offset+1] & 0xff)) << 16) |
		((int(sum[offset+2] & 0xff)) << 8) |
		(int(sum[offset+3]) & 0xff))

	f := fmt.Sprintf("%%0%dd", g.amount)
	mod := int32(value % int64(math.Pow10(g.amount)))

	token := fmt.Sprintf(f, mod)
	g.lastToken = token
	return token, nil
}
