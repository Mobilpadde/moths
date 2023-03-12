package option

import (
	"github.com/Mobilpadde/moths/v6/token"
)

// OptionFromEncoded is used as
// a shortcut from creating a new
// generator and importing the
// encoded string
func OptionFromEncoded(str string) token.Option {
	return func(g *token.Generator) error {
		return g.Import(str)
	}
}
