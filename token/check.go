package token

import (
	"github.com/Mobilpadde/moths/v6/token/checks"
)

// Check is checking if everything
// is working
//
// Meaning no errors is generated
// from getting codes
func (g *Generator) Check() error {
	if err := checks.CheckAmount(g.amount); err != nil {
		return err
	}

	if err := checks.CheckEmojies(g.emojies); err != nil {
		return err
	}

	if err := checks.CheckPeriod(g.period); err != nil {
		return err
	}

	if _, err := g.getToken(); err != nil {
		return err
	}

	return nil
}
