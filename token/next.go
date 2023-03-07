package token

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math"
	"time"

	"github.com/Mobilpadde/moths/v4/token/otp"
)

func (m *Moths) Next() (otp.OTP, error) {
	token, err := m.getToken()
	if err != nil {
		return otp.OTP{}, err
	}

	return otp.NewOTP(token, m.amount, m.emojies)
}

func (m *Moths) getToken() (string, error) {
	m.timing.curr = time.Now().UTC()

	since := m.timing.curr.Sub(m.timing.last)
	if since < m.interval {
		return m.lastToken, nil
	}

	m.timing.last = m.timing.curr
	m.timing.time = m.timing.time.Add(since)

	interval := uint64(m.timing.time.Unix() / int64(m.interval.Seconds()))

	// https://github.com/pquerna/otp/blob/master/hotp/hotp.go#L95-L123
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
