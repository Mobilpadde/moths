package checks_test

import (
	"errors"
	"testing"

	"github.com/Mobilpadde/moths/v6/token/checks"
	"github.com/Mobilpadde/moths/v6/token/errs"
)

func TestCheckAmount(t *testing.T) {
	if err := checks.CheckAmount(1); err != nil {
		t.Error("Expected to not return an error when checking ok. amount", err)
	}
}

func TestCheckAmountError(t *testing.T) {
	err := checks.CheckAmount(0)

	if !errors.Is(err, errs.ErrAmount1) {
		t.Error("Expected to return an error when checking not ok. amount:", err)
	}
}
