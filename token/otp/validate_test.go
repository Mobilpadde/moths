package otp

import (
	"testing"

	"github.com/Mobilpadde/moths/v4/token/emojies"
)

func TestValidate(t *testing.T) {
	amount := 6
	token := "000000"
	correct := "😸😼😹😺😸😻"

	otp, err := NewOTP(token, amount, emojies.CATS)
	if err != nil {
		t.Error("Expected to not return an error when creating new OTP:", err)
	}

	if !otp.Validate(correct) {
		t.Error("Expected code to be", correct, "not", otp.emojies)
	}
}

func TestValidateToken(t *testing.T) {
	amount := 6
	token := "000000"

	otp, err := NewOTP(token, amount, emojies.CATS)
	if err != nil {
		t.Error("Expected to not return an error when creating new OTP:", err)
	}

	if !otp.ValidateToken(token) {
		t.Error("Expected code to be", token, "not", otp.token)
	}
}
