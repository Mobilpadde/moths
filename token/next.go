package token

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math"
	"time"

	"github.com/Mobilpadde/moths/v4/token/code"
)

func (m *Generator) Next() (code.Code, error) {
	token, err := m.getToken()
	if err != nil {
		return code.Code{}, err
	}

	return code.NewCode(token, m.amount, m.emojies, m.timing.time, m.timing.time.Add(m.interval))
}

func (m *Generator) getToken() (string, error) {
	m.timing.curr = time.Now().UTC()

	since := m.timing.curr.Sub(m.timing.last)
	if since < m.interval {
		return m.lastToken, nil
	}

	m.timing.last = m.timing.curr
	m.timing.time = m.timing.time.Add(since)

	interval := uint64(m.timing.time.Unix() / int64(m.interval.Seconds()))

	// https://github.com/pquerna/code/blob/master/hotp/hotp.go#L95-L123
	buf := make([]byte, 8)
	h := hmac.New(sha1.New, m.secret)
	binary.BigEndian.PutUint64(buf, interval)

	h.Write(buf)
	sum := h.Sum(nil)

	// "Dynamic truncation" in RFC 4226
	// http://tools.ietf.org/html/rfc4226#section-5.4
	offset := sum[len(sum)-1] & 0xf
	value := int64(((int(sum[offset]) & 0x7f) << 24) |
		((int(sum[offset+1] & 0xff)) << 16) |
		((int(sum[offset+2] & 0xff)) << 8) |
		(int(sum[offset+3]) & 0xff))

	f := fmt.Sprintf("%%0%dd", m.amount)
	mod := int32(value % int64(math.Pow10(m.amount)))

	token := fmt.Sprintf(f, mod)
	m.lastToken = token
	return token, nil
}
