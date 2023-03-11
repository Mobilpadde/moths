package emojies

import "strings"

type Emojies []Emoji

func (emojies Emojies) String() string {
	return string(emojies)
}

func (emojies Emojies) SpacedString() string {
	return strings.Join(emojies.Slice(), " ")
}

func (emojies Emojies) Slice() []string {
	str := emojies.String()
	return strings.Split(str, "")
}
