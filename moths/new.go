package moths

import (
	"time"

	"github.com/Mobilpadde/moths/moths/checks"
)

func NewMoths(opts ...option) (*Moths, error) {
	m := &Moths{
		interval: 0,
		amount:   0,
		time:     time.Now(),
	}

	for _, opt := range opts {
		if err := opt(m); err != nil {
			return nil, err
		}
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

	if err := checks.CheckAmount(m.amount); err != nil {
		return nil, err
	}

	// check if everything is working
	_, err := m.getToken(true)

	// go m.update()
	return m, err
}
