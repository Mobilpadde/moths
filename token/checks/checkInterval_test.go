package checks

import (
	"errors"
	"testing"
	"time"

	"github.com/Mobilpadde/moths/v4/token/errs"
)

func TestCheckInterval(t *testing.T) {
	if err := CheckInterval(time.Second); err != nil {
		t.Error("Expected to not return an error when checking ok. interval:", err)
	}
}

func TestCheckIntervalError(t *testing.T) {
	err := CheckInterval(0)

	if !errors.Is(err, errs.ErrInterval1) {
		t.Error("Expected to return an error when checking not ok. interval:", err)
	}
}
