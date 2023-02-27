package otp

import (
	"math/rand"
	"strconv"
)

func bufferFromToken(token string) func(step int) ([]byte, error) {
	// // Does not support any more than a 15-characters
	// buf := []byte(token)
	// seed := int64(binary.LittleEndian.Uint64(buf))

	// Works for longer tokens
	seed, _ := strconv.ParseInt(token, 10, 64)
	r := rand.New(rand.NewSource(seed))

	return func(step int) ([]byte, error) {
		buffer := make([]byte, step)
		if _, err := r.Read(buffer); err != nil {
			return nil, err
		}

		return buffer, nil
	}
}
