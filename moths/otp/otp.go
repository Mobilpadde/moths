package otp

import (
	"math"
	"math/bits"
	"strings"

	"moths/moths/emojies"
)

const emojiBytes = 4

func NewOTP(token string, amount int, em emojies.Emojies) (OTP, error) {
	emojiAmount := len(em)
	size := emojiBytes * amount

	// https://github.com/tilaklodha/google-authenticator
	mask := 2<<uint32(31-bits.LeadingZeros32(uint32(emojiAmount-1|1))) - 1
	step := int(math.Ceil(1.6 * float64(mask*size) / float64(emojiAmount)))

	moth := new(strings.Builder)
	moth.Grow(size)

	buffer := bufferFromToken(token)
loop:
	for {
		randomBuffer, err := buffer(step)
		if err != nil {
			return OTP{}, err
		}

		for i := 0; i < step; i++ {
			currentIndex := int(randomBuffer[i]) & mask

			if currentIndex < emojiAmount {
				if _, err := moth.WriteRune(em[currentIndex]); err != nil {
					return OTP{}, err
				} else if moth.Len() == size {
					break loop
				}

				if moth.Len() > size {
					moth.Reset()
				}
			}
		}
	}

	split := strings.Split(moth.String(), "")
	return OTP{
		token:   token,
		emojies: emojies.ToEmojies(split),
	}, nil
}
