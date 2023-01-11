package moths

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/Mobilpadde/moths/moths/emojies"
	"github.com/Mobilpadde/moths/moths/errs"
)

func TestNewMoths(t *testing.T) {
	amount := 6
	secret := strings.Repeat("a", 32)

	if _, err := NewMoths(
		WithSecret(secret),
		WithInterval(time.Second),
		WithAmount(amount),
		WithEmojies(emojies.CATS),
	); err != nil {
		t.Error("Expected to not return an error when creating new Moths:", err)
	}
}

func TestNewMothsNoSecret(t *testing.T) {
	amount := 6

	_, err := NewMoths(
		WithInterval(time.Second),
		WithAmount(amount),
		WithEmojies(emojies.CATS),
	)

	if !errors.Is(err, errs.ErrKey20) {
		t.Error("Expected to return an error when creating new Moths without a secret:", err)
	}
}

func TestNewMothsNoInterval(t *testing.T) {
	amount := 6
	secret := strings.Repeat("a", 32)

	_, err := NewMoths(
		WithSecret(secret),
		WithAmount(amount),
		WithEmojies(emojies.CATS),
	)

	if !errors.Is(err, errs.ErrInterval1) {
		t.Error("Expected to return an error when creating new Moths without an interval:", err)
	}
}

func TestNewMothsNoAmount(t *testing.T) {
	secret := strings.Repeat("a", 32)

	_, err := NewMoths(
		WithSecret(secret),
		WithInterval(time.Second),
		WithEmojies(emojies.CATS),
	)

	if !errors.Is(err, errs.ErrAmount1) {
		t.Error("Expected to return an error when creating new Moths without an amount:", err)
	}
}

func TestNewMothsNoEmojies(t *testing.T) {
	amount := 6
	secret := strings.Repeat("a", 32)

	_, err := NewMoths(
		WithSecret(secret),
		WithInterval(time.Second),
		WithAmount(amount),
	)

	if !errors.Is(err, errs.ErrEmojies1) {
		t.Error("Expected to return an error when creating new Moths without any emojies:", err)
	}
}
