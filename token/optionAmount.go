package token

import (
	"github.com/Mobilpadde/moths/v5/token/checks"
)

func OptionWithAmount(amount int) Option {
	return func(m *Generator) error {
		if err := checks.CheckAmount(amount); err != nil {
			return err
		}

		m.amount = amount
		return nil
	}
}
