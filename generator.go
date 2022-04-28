package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/bits"
	"math/rand"
	"strings"
)

// randomBufferFromHOTP generates a random buffer from a HOTP token.
//
// /!\ Modified /!\
//
// Proudly stolen from https://stackoverflow.com/a/48307199/754471
func randomBufferFromHOTP(otp string) func(step int) ([]byte, error) {
	buf := []byte(fmt.Sprintf("%08s", otp))
	seed := int64(binary.BigEndian.Uint64(buf))
	r := rand.New(rand.NewSource(seed))

	return func(step int) ([]byte, error) {
		buffer := make([]byte, step)
		if _, err := r.Read(buffer); err != nil {
			return nil, err
		}
		return buffer, nil
	}
}

// MothsGenerator is essentially the nanoid generator with a few modifications.
//
// /!\ Modified /!\
//
// https://github.com/aidarkhanov/nanoid/blob/master/nanoid.go
func MothsGenerator(token string, numEmojies int, alphabet ...string) (string, error) {
	emojiBytes := 4
	size := emojiBytes * numEmojies

	mask := 2<<uint32(31-bits.LeadingZeros32(uint32(len(alphabet)-1|1))) - 1
	step := int(math.Ceil(1.6 * float64(mask*size) / float64(len(alphabet))))

	id := new(strings.Builder)
	id.Grow(size)

	buffer := randomBufferFromHOTP(token)
	for {
		randomBuffer, err := buffer(step)
		if err != nil {
			return "", err
		}

		for i := 0; i < step; i++ {
			currentIndex := int(randomBuffer[i]) & mask

			if currentIndex < len(alphabet) {
				if _, err := id.WriteString(alphabet[currentIndex]); err != nil {
					return "", err
				} else if id.Len() == size {
					return id.String(), nil
				}
			}
		}
	}
}
