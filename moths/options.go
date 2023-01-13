package moths

import (
	"encoding/base32"
	"strings"
	"time"

	"github.com/Mobilpadde/moths/moths/checks"
	"github.com/Mobilpadde/moths/moths/emojies"
)

type Moths struct {
	secret   []byte
	interval time.Duration
	amount   int
	emojies  emojies.Emojies
	time     time.Time
}

type option func(*Moths) error

func OptionWithSecret(secret string) option {
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

func OptionWithInterval(interval time.Duration) option {
	return func(m *Moths) error {
		if err := checks.CheckInterval(interval); err != nil {
			return err
		}

		m.interval = interval
		return nil
	}
}

func OptionWithAmount(amount int) option {
	return func(m *Moths) error {
		if err := checks.CheckAmount(amount); err != nil {
			return err
		}

		m.amount = amount
		return nil
	}
}

func OptionWithEmojies(emojies emojies.Emojies) option {
	return func(m *Moths) error {
		if err := checks.CheckEmojies(emojies); err != nil {
			return err
		}

		m.emojies = emojies
		return nil
	}
}

func OptionWithTime(t time.Time) option {
	return func(m *Moths) error {
		m.time = t
		return nil
	}
}
