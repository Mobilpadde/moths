package token

import (
	"time"
)

// NewGenerator creates a new token generator
// with the given [Option]s.
func NewGenerator(opts ...Option) (*Generator, error) {
	g := &Generator{
		amount: 6, // Defaults to `6` as most other TOTP codes are this length
	}

	for _, opt := range opts {
		if err := opt(g); err != nil {
			return nil, err
		}
	}

	now := time.Now().UTC()
	g.timing.curr = now
	g.timing.last = now.Add(-g.period)

	return g, nil
}
