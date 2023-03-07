package token

import "time"

// OptionWithTime is used to specify
// a custom time to generate code from.
//
// If none is specified, the current time
// will be used.
func OptionWithTime(t time.Time) Option {
	return func(m *Generator) error {
		m.timing.time = t
		return nil
	}
}
