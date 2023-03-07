package code

import (
	"testing"
	"time"

	"github.com/Mobilpadde/moths/v4/token/emojies"
)

func TestNewCode(t *testing.T) {
	amount := 6
	token := "000000"
	correct := "😸😼😹😺😸😻"

	code, err := NewCode(token, amount, emojies.CATS, time.Time{}, time.Time{})
	if err != nil {
		t.Error("Expected to not return an error when creating new code:", err)
	}

	if code.Token() != token {
		t.Error("Expected token to be", token, "not", code.Token())
	}

	if code.String() != correct {
		t.Error("Expected code to be", correct, "not", code.String())
	}
}

func TestCreatedAt(t *testing.T) {
	amount := 6
	token := "000000"
	createdAt := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)

	code, err := NewCode(token, amount, emojies.CATS, createdAt, time.Time{})
	if err != nil {
		t.Error("Expected to not return an error when creating new code:", err)
	}

	if code.CreatedAt().Compare(createdAt) != 0 {
		t.Error("Expected to be created at", createdAt, "not", code.CreatedAt())
	}
}

func TestExpiresAt(t *testing.T) {
	amount := 6
	token := "000000"
	expiresAt := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)

	code, err := NewCode(token, amount, emojies.CATS, time.Time{}, expiresAt)
	if err != nil {
		t.Error("Expected to not return an error when creating new code:", err)
	}

	if code.ExpiresAt().Compare(expiresAt) != 0 {
		t.Error("Expected to expire at", expiresAt, "not", code.ExpiresAt())
	}
}
