package token

import (
	"github.com/Mobilpadde/moths/v4/token/checks"
	"github.com/Mobilpadde/moths/v4/token/emojies"
)

func OptionWithEmojies(emojies emojies.Emojies) Option {
	return func(m *Generator) error {
		if err := checks.CheckEmojies(emojies); err != nil {
			return err
		}

		m.emojies = emojies
		return nil
	}
}
