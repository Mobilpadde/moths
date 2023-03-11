package code_test

import (
	"testing"
	"time"

	"github.com/Mobilpadde/moths/v5/token/code"
	"github.com/Mobilpadde/moths/v5/token/emojies"
)

func TestValidate(t *testing.T) {
	amount := 6
	token := "000000"
	correct := "😸😼😹😺😸😻"

	code, err := code.NewCode(token, amount, emojies.CATS, time.Time{}, time.Time{})
	if err != nil {
		t.Error("Expected to not return an error when creating new code:", err)
	}

	if !code.Validate(correct) {
		t.Error("Expected code to be", correct, "not", code.String())
	}
}

func TestValidateToken(t *testing.T) {
	amount := 6
	token := "000000"

	code, err := code.NewCode(token, amount, emojies.CATS, time.Time{}, time.Time{})
	if err != nil {
		t.Error("Expected to not return an error when creating new code:", err)
	}

	if !code.ValidateToken(token) {
		t.Error("Expected code to be", token, "not", code.Token())
	}
}
