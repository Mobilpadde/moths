package token

import (
	"time"

	"github.com/Mobilpadde/moths/v5/token/checks"
)

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
	if m.timing.time == (time.Time{}) {
		m.timing.time = now
	}

	if err := checks.CheckPeriod(m.period); err != nil {
		return nil, err
	}

	if err := checks.CheckEmojies(m.emojies); err != nil {
		return nil, err
	}

	// Check if everything is working
	if _, err := m.getToken(); err != nil {
		return nil, err
	}

	return m, nil
}
