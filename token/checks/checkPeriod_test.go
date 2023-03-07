package checks

import (
	"errors"
	"testing"
	"time"

	"github.com/Mobilpadde/moths/v4/token/errs"
)

func TestCheckInterval(t *testing.T) {
	if err := CheckPeriod(time.Second); err != nil {
		t.Error("Expected to not return an error when checking ok. interval:", err)
	}
}

func TestCheckIntervalError(t *testing.T) {
	err := CheckPeriod(0)

	if !errors.Is(err, errs.ErrPeriod1) {
		t.Error("Expected to return an error when checking not ok. interval:", err)
	}
}
