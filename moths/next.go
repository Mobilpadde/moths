package moths

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"time"

	"moths/moths/otp"
)

func (m *Moths) Next() (otp.OTP, error) {
	token, err := m.getToken(true)
	if err != nil {
		return otp.OTP{}, err
	}

	return otp.NewOTP(token, m.emojies)
}

func (m *Moths) getToken(bump bool) (string, error) {
	if bump {
		// Bump the time on every run
		m.time = m.time.Add(time.Since(m.time))
	}

	b := make([]byte, 8)
	interval := uint64(m.time.Unix() / int64(m.interval))
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
	return fmt.Sprintf("%06d", otp), nil
}
