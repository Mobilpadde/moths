package checks_test

import (
	"errors"
	"testing"

	"github.com/Mobilpadde/moths/v6/token/checks"
	"github.com/Mobilpadde/moths/v6/token/emojies"
	"github.com/Mobilpadde/moths/v6/token/errs"
	"github.com/enescakir/emoji"
)

func TestCheckEmojies(t *testing.T) {
	emoji := emoji.BlackCat.String()
	emojies := emojies.ToEmojies([]string{emoji})

	if err := checks.CheckEmojies(emojies); err != nil {
		t.Error("Expected to not return an error when checking ok. emoji length:", err)
	}
}

func TestCheckEmojiesError(t *testing.T) {
	emojies := emojies.ToEmojies([]string{})
	err := checks.CheckEmojies(emojies)

	if !errors.Is(err, errs.ErrEmojies1) {
		t.Error("Expected to return an error when checking not ok. emoji length:", err)
	}
}
