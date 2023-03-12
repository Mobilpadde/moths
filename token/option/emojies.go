package option

import (
	"github.com/Mobilpadde/moths/v6/token"
	"github.com/Mobilpadde/moths/v6/token/emojies"
)

// OptionWithEmojies is used to specify
// which emojies that can be used in
// any given code.
func OptionWithEmojies(emojies emojies.Emojies) token.Option {
	return func(g *token.Generator) error {
		return g.SetEmojies(emojies)
	}
}
