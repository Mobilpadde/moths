package token

import (
	"time"
)

// NewGenerator creates a new token generator
// with the given [Option]s.
func NewGenerator(opts ...Option) (*Generator, error) {
	m := &Generator{
		amount: 6, // Defaults to `6` as most other TOTP codes are this length
	}

	for _, opt := range opts {
		if err := opt(m); err != nil {
			return nil, err
		}
	}

	now := time.Now().UTC()
	m.timing.curr = now
	m.timing.last = now.Add(-m.period)

	return m, nil
}
