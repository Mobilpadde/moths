package checks

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/Mobilpadde/moths/moths/emojies"
	"github.com/Mobilpadde/moths/moths/errs"
	"github.com/enescakir/emoji"
)

func TestCheckSecret(t *testing.T) {
	secret := strings.Repeat("a", 32)
	if err := CheckSecret(secret); err != nil {
		t.Error("Expected to not return an error when checking ok secret:", err)
	}
}

func TestCheckSecretError(t *testing.T) {
	secret := "0"
	err := CheckSecret(secret)

	if !errors.Is(err, errs.ErrSecret32) {
		t.Error("Expected to return an error when checking not ok secret:", err)
	}
}

func TestCheckSecretKey(t *testing.T) {
	secret := []byte(strings.Repeat("a", 20))
	if err := CheckSecretKey(secret); err != nil {
		t.Error("Expected to not return an error when checking ok secret key:", err)
	}
}

func TestCheckSecretKeyError(t *testing.T) {
	secret := []byte{0}
	err := CheckSecretKey(secret)

	if !errors.Is(err, errs.ErrKey20) {
		t.Error("Expected to return an error when checking not ok secret key:", err)
	}
}

func TestCheckEmojies(t *testing.T) {
	emoji := emoji.BlackCat.String()
	emojies := emojies.ToEmojies([]string{emoji})

	if err := CheckEmojies(emojies); err != nil {
		t.Error("Expected to not return an error when checking ok emoji length:", err)
	}
}

func TestCheckEmojiesError(t *testing.T) {
	emojies := emojies.ToEmojies([]string{})
	err := CheckEmojies(emojies)

	if !errors.Is(err, errs.ErrEmojies1) {
		t.Error("Expected to return an error when checking not ok emoji length:", err)
	}
}

func TestCheckInterval(t *testing.T) {
	if err := CheckInterval(time.Second); err != nil {
		t.Error("Expected to not return an error when checking ok interval:", err)
	}
}

func TestCheckIntervalError(t *testing.T) {
	err := CheckInterval(0)

	if !errors.Is(err, errs.ErrInterval1) {
		t.Error("Expected to return an error when checking not ok interval:", err)
	}
}

func TestCheckAmount(t *testing.T) {
	if err := CheckAmount(1); err != nil {
		t.Error("Expected to not return an error when checking ok amount", err)
	}
}

func TestCheckAmountError(t *testing.T) {
	err := CheckAmount(0)

	if !errors.Is(err, errs.ErrAmount1) {
		t.Error("Expected to return an error when checking not ok amount:", err)
	}
}
