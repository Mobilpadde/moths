package emojies

import (
	"strings"
)

type Emojies []string

func (emojies Emojies) String() string {
	return strings.Join(emojies, "")
}

func (emojies Emojies) SpacedString() string {
	return strings.Join(emojies, " ")
}

func (emojies Emojies) Slice() []string {
	return emojies
}
