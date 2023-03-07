package token

import (
	"time"

	"github.com/Mobilpadde/moths/v5/token/checks"
)

// OptionWithPeriod is used to specify
// how long a code will remain valid.
func OptionWithPeriod(period time.Duration) Option {
	return func(m *Generator) error {
		if err := checks.CheckPeriod(period); err != nil {
			return err
		}

		m.period = period
		return nil
	}
}
