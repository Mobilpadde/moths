package option

import (
	"time"

	"github.com/Mobilpadde/moths/v5/token"
)

// OptionWithPeriod is used to specify
// how long a code will remain valid.
func OptionWithPeriod(period time.Duration) token.Option {
	return func(g *token.Generator) error {
		return g.SetPeriod(period)
	}
}
