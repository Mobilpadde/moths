package checks

import (
	"github.com/Mobilpadde/moths/v4/token/errs"
)

func CheckAmount(amount int) error {
	if amount < 1 {
		return errs.ErrAmount1
	}

	return nil
}
