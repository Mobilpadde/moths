package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"log"
	"strconv"
)

// GetHOTPToken returns a HOTP token for the given secret and interval.
//
// /!\ Modified /!\
//
// https://github.com/tilaklodha/google-authenticator/blob/master/google_authenticator.go
func GetHOTPToken(secret string, interval int64) string {
	key := []byte(fmt.Sprintf("%08s", secret))

	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(interval))

	// Signing the value using HMAC-SHA1 Algorithm
	hash := hmac.New(sha1.New, key)
	hash.Write(bs)
	h := hash.Sum(nil)

	// We're going to use a subset of the generated hash.
	// Using the last nibble (half-byte) to choose the index to start from.
	// This number is always appropriate as it's maximum decimal 15, the hash will
	// have the maximum index 19 (20 bytes of SHA1) and we need 4 bytes.
	o := (h[19] & 15)

	var header uint32
	// Get 32 bit chunk from hash starting at the o
	r := bytes.NewReader(h[o : o+4])
	err := binary.Read(r, binary.BigEndian, &header)
	if err != nil {
		log.Fatalln(err)
	}

	// Ignore most significant bits as per RFC 4226.
	// Takes division from one million to generate a remainder less than < 7 digits
	h12 := (int(header) & 0x7fffffff) % 1000000

	// Converts number as a string
	otp := strconv.Itoa(int(h12))

	// Padding with 0s to the left if the OTP is less than 8 digits
	return fmt.Sprintf("%06s", otp)
}
