package checks

import (
	"time"

	"github.com/Mobilpadde/moths/v5/token/errs"
)

func CheckPeriod(period time.Duration) error {
	if period.Milliseconds() < 1000 {
		return errs.ErrPeriod1
	}

	return nil
}
