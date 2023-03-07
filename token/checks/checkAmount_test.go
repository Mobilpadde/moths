package checks

import (
	"errors"
	"testing"

	"github.com/Mobilpadde/moths/v5/token/errs"
)

func TestCheckAmount(t *testing.T) {
	if err := CheckAmount(1); err != nil {
		t.Error("Expected to not return an error when checking ok. amount", err)
	}
}

func TestCheckAmountError(t *testing.T) {
	err := CheckAmount(0)

	if !errors.Is(err, errs.ErrAmount1) {
		t.Error("Expected to return an error when checking not ok. amount:", err)
	}
}
