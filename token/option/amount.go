package option

import (
	"github.com/Mobilpadde/moths/v5/token"
)

// OptionWithAmount is used to specify
// the amount of emojies in a code.
func OptionWithAmount(amount int) token.Option {
	return func(g *token.Generator) error {
		return g.SetAmount(amount)
	}
}
