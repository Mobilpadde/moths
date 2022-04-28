package moths

import (
	"errors"
	"math"
	"math/bits"
	"strings"
)

const emojiBytes = 4

func (m *Moths) Next(numEmojies int, alphabet []string) (string, error) {
	if len(alphabet) < 1 {
		return "", errors.New("alphabet must be at least 1 character")
	}

	size := emojiBytes * numEmojies
	// https://github.com/tilaklodha/google-authenticator
	mask := 2<<uint32(31-bits.LeadingZeros32(uint32(len(alphabet)-1|1))) - 1
	step := int(math.Ceil(1.6 * float64(mask*size) / float64(len(alphabet))))

	id := new(strings.Builder)
	id.Grow(size)

	tkn, err := m.getHOTPToken()
	if err != nil {
		return "", err
	}

	buffer := randomBufferFromHOTP(tkn)
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
