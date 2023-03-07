package token

import (
	"github.com/Mobilpadde/moths/v5/token/checks"
	"github.com/Mobilpadde/moths/v5/token/emojies"
)

// OptionWithEmojies is used to specify
// which emojies that can be used in
// any given code.
func OptionWithEmojies(emojies emojies.Emojies) Option {
	return func(m *Generator) error {
		if err := checks.CheckEmojies(emojies); err != nil {
			return err
		}

		m.emojies = emojies
		return nil
	}
}
