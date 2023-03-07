package moths

import (
	"time"

	"github.com/Mobilpadde/moths/v4/moths/checks"
)

func NewMoths(opts ...option) (*Moths, error) {
	m := &Moths{
		interval: 0,
		amount:   6, // Defaults to `6` as most other TOTP services uses that
	}

	for _, opt := range opts {
		if err := opt(m); err != nil {
			return nil, err
		}
	}

	now := time.Now().UTC()
	m.timing.curr = now
	m.timing.last = now.Add(-m.interval)
	if m.timing.time == (time.Time{}) {
		m.timing.time = now
	}

	if err := checks.CheckSecretKey(m.secret); err != nil {
		return nil, err
	}

	if err := checks.CheckInterval(m.interval); err != nil {
		return nil, err
	}

	if err := checks.CheckEmojies(m.emojies); err != nil {
		return nil, err
	}

	// Checks if everything is working
	if _, err := m.getToken(); err != nil {
		return nil, err
	}

	return m, nil
}
