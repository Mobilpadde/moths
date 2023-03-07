package moths

import (
	"github.com/Mobilpadde/moths/v4/moths/otp"
)

func (m *Moths) Validate(moth string) bool {
	if len(moth)/otp.EmojiBytes != m.amount {
		return false
	}

	token, err := m.getToken()
	if err != nil {
		return false
	}

	same, err := otp.NewOTP(token, m.amount, m.emojies)
	if err != nil {
		return false
	}

	return same.Validate(moth)
}

// This should maybe not be used
// as you should not really expose the `token`
// to your users
func (m *Moths) ValidateToken(oldToken string) bool {
	token, err := m.getToken()
	if err != nil {
		return false
	}

	same, err := otp.NewOTP(token, m.amount, m.emojies)
	if err != nil {
		return false
	}

	return same.ValidateToken(oldToken)
}
