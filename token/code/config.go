package code

import (
	"github.com/Mobilpadde/moths/v4/token/emojies"
)

type Code struct {
	token   string
	emojies emojies.Emojies
}

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
