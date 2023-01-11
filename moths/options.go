package moths

import (
	"encoding/base32"
	"strings"
	"time"

	"moths/moths/checks"
	"moths/moths/emojies"
)

type Moths struct {
	secret   []byte
	interval time.Duration
	amount   int
	emojies  emojies.Emojies
	time     time.Time
}

type option func(*Moths) error

func WithSecret(secret string) option {
	return func(m *Moths) error {
		if err := checks.CheckSecret(secret); err != nil {
			return err
		}

		// Setup the secret as an encoded key
		secret = strings.ToUpper(secret)
		key, err := base32.StdEncoding.DecodeString(secret)
		if err != nil {
			return err
		}

		m.secret = key
		return nil
	}
}

func WithInterval(interval time.Duration) option {
	return func(m *Moths) error {
		if err := checks.CheckInterval(interval); err != nil {
			return err
		}

		m.interval = interval
		return nil
	}
}

func WithAmount(amount int) option {
	return func(m *Moths) error {
		// TODO: Check if amount is okay (>0?)
		m.amount = amount
		return nil
	}
}

func WithEmojies(emojies emojies.Emojies) option {
	return func(m *Moths) error {
		if err := checks.CheckEmojies(emojies); err != nil {
			return err
		}

		m.emojies = emojies
		return nil
	}
}

func WithTime(t time.Time) option {
	return func(m *Moths) error {
		m.time = t
		return nil
	}
}
