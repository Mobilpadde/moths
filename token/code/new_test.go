package code

import (
	"testing"

	"github.com/Mobilpadde/moths/v4/token/emojies"
)

func TestNewCode(t *testing.T) {
	amount := 6
	token := "000000"
	correct := "😸😼😹😺😸😻"

	code, err := NewCode(token, amount, emojies.CATS)
	if err != nil {
		t.Error("Expected to not return an error when creating new code:", err)
	}

	if code.token != token {
		t.Error("Expected token to be", token, "not", code.token)
	}

	if code.emojies.String() != correct {
		t.Error("Expected code to be", correct, "not", code.emojies)
	}
}
