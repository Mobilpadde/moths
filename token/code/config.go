package code

import (
	"time"

	"github.com/Mobilpadde/moths/v4/token/emojies"
)

type Code struct {
	token     string
	emojies   emojies.Emojies
	createdAt time.Time
	expiresAt time.Time
}

// Deprecated: Insecure. Use `String`, `SpacedString` or `Slice` instead.
func (code Code) Token() string {
	return code.token
}

func (code Code) String() string {
	return code.emojies.String()
}

func (code Code) SpacedString() string {
	return code.emojies.SpacedString()
}

func (code Code) Slice() []string {
	return code.emojies.Slice()
}

func (code Code) CreatedAt() time.Time {
	return code.createdAt
}

func (code Code) ExpiresAt() time.Time {
	return code.expiresAt
}
