package code

import (
	"math"
	"math/bits"
	"strings"
	"time"

	"github.com/Mobilpadde/moths/v5/token/emojies"
)

const EmojiBytes = 4

// NewCode creates a new code
//
// This is generally not for public usage.
func NewCode(token string, amount int, em emojies.Emojies, createdAt time.Time, expiresAt time.Time) (Code, error) {
	code, err := generateCode(token, amount, em)
	if err != nil {
		return Code{}, err
	}

	return Code{
		token:     token,
		emojies:   code,
		createdAt: createdAt,
		expiresAt: expiresAt,
	}, nil
}

func generateCode(token string, amount int, em emojies.Emojies) (emojies.Emojies, error) {
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
			return emojies.EMPTY, err
		}

		for i := 0; i < step; i++ {
			currentIndex := int(randomBuffer[i]) & mask

			if currentIndex < emojiAmount {
				if _, err := code.WriteRune(em[currentIndex]); err != nil {
					return emojies.EMPTY, err
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
	return emojies.ToEmojies(split), nil
}
