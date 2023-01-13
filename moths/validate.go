package moths

import (
	"time"

	"github.com/Mobilpadde/moths/moths/otp"
)

func (m *Moths) Validate(moth string) bool {
	if len(moth)/otp.EmojiBytes != m.amount {
		return false
	}

	token, err := m.getToken(false)
	if err != nil {
		return false
	}

	same, err := otp.NewOTP(token, m.amount, m.emojies)
	if err != nil {
		return false
	}

	// Bump the time
	// m.getToken(true)
	m.time = m.time.Add(time.Since(m.time))

	return same.Validate(moth)
}

func (m *Moths) ValidateToken(code string) bool {
	token, err := m.getToken(false)
	if err != nil {
		return false
	}

	return token == code
}
