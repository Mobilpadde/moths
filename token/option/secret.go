package option

import "github.com/Mobilpadde/moths/v6/token"

// OptionWithSecret is used to specify
// the secret to generate codes from.
func OptionWithSecret(secret string) token.Option {
	return func(g *token.Generator) error {
		return g.SetSecret(secret)
	}
}
