package token

import (
	"time"

	"github.com/Mobilpadde/moths/v4/token/checks"
)

func OptionWithInterval(interval time.Duration) Option {
	return func(m *Generator) error {
		if err := checks.CheckInterval(interval); err != nil {
			return err
		}

		m.interval = interval
		return nil
	}
}
