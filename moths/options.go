package moths

import (
	"encoding/base64"
	"time"

	"github.com/Mobilpadde/moths/v4/moths/checks"
	"github.com/Mobilpadde/moths/v4/moths/emojies"
)

type Moths struct {
	secret  []byte
	amount  int
	emojies emojies.Emojies

	interval  time.Duration
	lastToken string

	timing struct {
		time time.Time
		curr time.Time
		last time.Time
	}
}

type option func(*Moths) error

func OptionWithSecret(secret string) option {
	return func(m *Moths) error {
		if err := checks.CheckSecret(secret); err != nil {
			return err
		}

		// Encode the secret as base64 - for extra security (🤷)
		data := []byte(secret)
		dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
		base64.StdEncoding.Encode(dst, data)

		m.secret = dst
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
		m.timing.time = t
		return nil
	}
}
