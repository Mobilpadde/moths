package checks_test

import (
	"errors"
	"testing"
	"time"

	"github.com/Mobilpadde/moths/v5/token/checks"
	"github.com/Mobilpadde/moths/v5/token/errs"
)

func TestCheckInterval(t *testing.T) {
	if err := checks.CheckPeriod(time.Second); err != nil {
		t.Error("Expected to not return an error when checking ok. interval:", err)
	}
}

func TestCheckIntervalError(t *testing.T) {
	err := checks.CheckPeriod(0)

	if !errors.Is(err, errs.ErrPeriod1) {
		t.Error("Expected to return an error when checking not ok. interval:", err)
	}
}
