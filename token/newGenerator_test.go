package token

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/Mobilpadde/moths/v4/token/emojies"
	"github.com/Mobilpadde/moths/v4/token/errs"
)

func TestNewGenerator(t *testing.T) {
	amount := 6
	secret := strings.Repeat("a", 32)

	if _, err := NewGenerator(
		OptionWithSecret(secret),
		OptionWithInterval(time.Second),
		OptionWithAmount(amount),
		OptionWithEmojies(emojies.CATS),
	); err != nil {
		t.Error("Expected to not return an error when creating new generator:", err)
	}
}

func TestNewGeneratorNoSecret(t *testing.T) {
	amount := 6

	_, err := NewGenerator(
		OptionWithSecret(""),
		OptionWithInterval(time.Second),
		OptionWithAmount(amount),
		OptionWithEmojies(emojies.CATS),
	)

	if !errors.Is(err, errs.ErrSecretLength) {
		t.Error("Expected to return an error when creating new generator without a secret:", err)
	}
}

func TestNewGeneratorNoInterval(t *testing.T) {
	amount := 6
	secret := strings.Repeat("a", 32)

	_, err := NewGenerator(
		OptionWithSecret(secret),
		OptionWithAmount(amount),
		OptionWithEmojies(emojies.CATS),
	)

	if !errors.Is(err, errs.ErrInterval1) {
		t.Error("Expected to return an error when creating new generator without an interval:", err)
	}
}

func TestNewGeneratorNoAmount(t *testing.T) {
	secret := strings.Repeat("a", 32)

	_, err := NewGenerator(
		OptionWithSecret(secret),
		OptionWithAmount(0),
		OptionWithInterval(time.Second),
		OptionWithEmojies(emojies.CATS),
	)

	if !errors.Is(err, errs.ErrAmount1) {
		t.Error("Expected to return an error when creating new generator without an amount:", err)
	}
}

func TestNewGeneratorNoEmojies(t *testing.T) {
	amount := 6
	secret := strings.Repeat("a", 32)

	_, err := NewGenerator(
		OptionWithSecret(secret),
		OptionWithInterval(time.Second),
		OptionWithAmount(amount),
	)

	if !errors.Is(err, errs.ErrEmojies1) {
		t.Error("Expected to return an error when creating new generator without any emojies:", err)
	}
}
