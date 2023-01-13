package moths

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/Mobilpadde/moths/moths/otp"
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

	b := make([]byte, 8)
	interval := uint64(m.timing.time.Unix() / int64(m.interval.Seconds()))
	binary.BigEndian.PutUint64(b, interval)

	hash := hmac.New(sha1.New, m.secret)
	hash.Write(b)
	h := hash.Sum(nil)

	o := (h[19] & 15)
	r := bytes.NewReader(h[o : o+4])

	var header uint32
	err := binary.Read(r, binary.BigEndian, &header)
	if err != nil {
		return "", err
	}

	otp := (int(header) & 0x7fffffff) % 1000000
	pad := fmt.Sprintf("%06d", otp)

	m.lastToken = pad
	return pad, nil
}
