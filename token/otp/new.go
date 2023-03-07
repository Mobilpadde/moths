package otp

import (
	"math"
	"math/bits"
	"strings"

	"github.com/Mobilpadde/moths/v4/token/emojies"
)

const EmojiBytes = 4

func NewOTP(token string, amount int, em emojies.Emojies) (OTP, error) {
	emojiAmount := len(em)
	size := EmojiBytes * amount

	code := strings.Builder{}
	code.Grow(size)
	buffer := bufferFromToken(token)

	// https://github.com/tilaklodha/google-authenticator
	mask := 2<<uint32(31-bits.LeadingZeros32(uint32(emojiAmount-1|1))) - 1
	step := int(math.Ceil(1.6 * float64(mask*size) / float64(emojiAmount)))

loop:
	for {
		randomBuffer, err := buffer(step)
		if err != nil {
			return OTP{}, err
		}

		for i := 0; i < step; i++ {
			currentIndex := int(randomBuffer[i]) & mask

			if currentIndex < emojiAmount {
				if _, err := code.WriteRune(em[currentIndex]); err != nil {
					return OTP{}, err
				} else if code.Len() == size {
					break loop
				}

				if code.Len() > size {
					code.Reset()
				}
			}
		}
	}

	split := strings.Split(code.String(), "")
	return OTP{
		token:   token,
		emojies: emojies.ToEmojies(split),
	}, nil
}
