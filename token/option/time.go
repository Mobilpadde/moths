package option

import (
	"time"

	"github.com/Mobilpadde/moths/v5/token"
)

// OptionWithTime is used to specify
// a custom time to generate code from.
//
// If none is specified, the current time
// will be used.
func OptionWithTime(t time.Time) token.Option {
	return func(g *token.Generator) error {
		return g.SetTime(t)
	}
}
