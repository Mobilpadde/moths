package token

import (
	"strings"
	"testing"
	"time"

	"github.com/Mobilpadde/moths/v4/token/emojies"
)

func TestNextAndValidate(t *testing.T) {
	amount := 6
	secret := strings.Repeat("a", 32)

	gen, err := NewMoths(
		OptionWithSecret(secret),
		OptionWithInterval(time.Second),
		OptionWithAmount(amount),
		OptionWithEmojies(emojies.CATS),
		OptionWithTime(time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)),
	)
	if err != nil {
		t.Error("Expected to not return an error when creating new Moths:", err)
	}

	otp, err := gen.Next()
	if err != nil {
		t.Error("Expected to not return an error when creating new OTP:", err)
	}

	errStr := "Expected %s to be %s not %s"
	correct := "🙀😾😹🙀😼😹"
	if !gen.Validate(correct) {
		t.Errorf(errStr, "code", correct, otp)
	}

	correct = "187603"
	if !gen.ValidateToken(correct) {
		t.Errorf(errStr, "token", correct, otp.Token())
	}
}
