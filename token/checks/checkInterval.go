package checks

import (
	"time"

	"github.com/Mobilpadde/moths/v4/token/errs"
)

func CheckInterval(interval time.Duration) error {
	if interval.Milliseconds() < 1000 {
		return errs.ErrInterval1
	}

	return nil
}
