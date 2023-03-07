package token

import (
	"time"

	"github.com/Mobilpadde/moths/v4/token/checks"
)

func OptionWithPeriod(period time.Duration) Option {
	return func(m *Generator) error {
		if err := checks.CheckPeriod(period); err != nil {
			return err
		}

		m.period = period
		return nil
	}
}
