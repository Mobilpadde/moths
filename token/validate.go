package token

import (
	"time"

	"github.com/Mobilpadde/moths/v6/token/code"
)

func (g *Generator) Validate(moth string) bool {
	if len(moth)/code.EmojiBytes != g.amount {
		return false
	}

	token, err := g.getToken()
	if err != nil {
		return false
	}

	same, err := code.NewCode(token, g.amount, g.emojies, time.Time{}, time.Time{})
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
func (g *Generator) ValidateToken(oldToken string) bool {
	token, err := g.getToken()
	if err != nil {
		return false
	}

	same, err := code.NewCode(token, g.amount, g.emojies, time.Time{}, time.Time{})
	if err != nil {
		return false
	}

	return same.ValidateToken(oldToken)
}
