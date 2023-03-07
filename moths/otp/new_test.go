package otp

import (
	"testing"

	"github.com/Mobilpadde/moths/v4/moths/emojies"
)

func TestNewOTP(t *testing.T) {
	amount := 6
	token := "000000"
	correct := "😸😼😹😺😸😻"

	otp, err := NewOTP(token, amount, emojies.CATS)
	if err != nil {
		t.Error("Expected to not return an error when creating new OTP:", err)
	}

	if otp.token != token {
		t.Error("Expected token to be", token, "not", otp.token)
	}

	if otp.emojies.String() != correct {
		t.Error("Expected code to be", correct, "not", otp.emojies)
	}
}
