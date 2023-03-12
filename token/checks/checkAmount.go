package checks

import (
	"github.com/Mobilpadde/moths/v6/token/errs"
)

// CheckAmount is used to check
// if the length of a token can be used.
func CheckAmount(amount int) error {
	if amount < 1 {
		return errs.ErrAmount1
	}

	return nil
}
