package checks

import (
	"time"

	"github.com/Mobilpadde/moths/v6/token/errs"
)

// CheckPeriod is used to check
// if the period of validity is long enough.
func CheckPeriod(period time.Duration) error {
	if period.Milliseconds() < 1000 {
		return errs.ErrPeriod1
	}

	return nil
}
