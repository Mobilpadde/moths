package moths

import (
	"encoding/binary"
	"fmt"
	"math/rand"
)

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
