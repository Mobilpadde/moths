package token

import (
	"github.com/Mobilpadde/moths/v4/token/code"
)

func (m *Generator) Validate(moth string) bool {
	if len(moth)/code.EmojiBytes != m.amount {
		return false
	}

	token, err := m.getToken()
	if err != nil {
		return false
	}

	same, err := code.NewCode(token, m.amount, m.emojies)
	if err != nil {
		return false
	}

	return same.Validate(moth)
}

// This should maybe not be used
// as you should not really expose the `token`
// to your users
//
// Deprecated: Insecure. Use `Validate` instead.
func (m *Generator) ValidateToken(oldToken string) bool {
	token, err := m.getToken()
	if err != nil {
		return false
	}

	same, err := code.NewCode(token, m.amount, m.emojies)
	if err != nil {
		return false
	}

	return same.ValidateToken(oldToken)
}
