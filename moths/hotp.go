package moths

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (m *Moths) getHOTPToken() (string, error) {
	ks := strings.ToUpper(m.secret)
	key, err := base32.StdEncoding.DecodeString(ks)
	if err != nil {
		return "", err
	}

	bs := make([]byte, 8)
	interval := uint64(time.Now().Unix() / int64(m.interval))
	binary.BigEndian.PutUint64(bs, interval)

	hash := hmac.New(sha1.New, key)
	hash.Write(bs)
	h := hash.Sum(nil)
	o := (h[19] & 15)

	var header uint32
	r := bytes.NewReader(h[o : o+4])
	err = binary.Read(r, binary.BigEndian, &header)
	if err != nil {
		return "", err
	}

	h12 := (int(header) & 0x7fffffff) % 1000000
	otp := strconv.Itoa(int(h12))
	return fmt.Sprintf("%06s", otp), nil
}
