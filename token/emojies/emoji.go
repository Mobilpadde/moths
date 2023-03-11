package emojies

type Emoji rune

func (emoji Emoji) String() string {
	return string(emoji)
}
