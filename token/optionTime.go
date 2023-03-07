package token

import "time"

func OptionWithTime(t time.Time) Option {
	return func(m *Generator) error {
		m.timing.time = t
		return nil
	}
}
