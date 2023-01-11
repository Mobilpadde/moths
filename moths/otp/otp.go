package otp

import (
	"math"
	"math/bits"
	"strings"

	"moths/moths/emojies"
)

const emojiBytes = 4

func NewOTP(token string, amount int, emojies emojies.Emojies) (OTP, error) {
	emojiAmount := len(emojies)
	size := emojiBytes * amount

	// https://github.com/tilaklodha/google-authenticator
	mask := 2<<uint32(31-bits.LeadingZeros32(uint32(emojiAmount-1|1))) - 1
	step := int(math.Ceil(1.6 * float64(mask*size) / float64(emojiAmount)))

	id := new(strings.Builder)
	id.Grow(size)

	buffer := bufferFromToken(token)
loop:
	for {
		randomBuffer, err := buffer(step)
		if err != nil {
			return OTP{}, err
		}

		// TODO: This goes on forever? - figure why!
		// Use emojies.GESTURES
		for i := 0; i < step; i++ {
			currentIndex := int(randomBuffer[i]) & mask

			if currentIndex < len(emojies) {
				if _, err := id.WriteString(emojies[currentIndex]); err != nil {
					return OTP{}, err
				} else if id.Len() == size {
					break loop
				}
			}
		}
	}

	emojies = strings.Split(id.String(), "")
	return OTP{
		token:   token,
		emojies: emojies,
	}, nil
}
