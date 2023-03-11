package token_test

import (
	"strings"
	"testing"
	"time"

	"github.com/Mobilpadde/moths/v5/token"
	"github.com/Mobilpadde/moths/v5/token/emojies"
	"github.com/Mobilpadde/moths/v5/token/option"
)

func TestNextAndValidate(t *testing.T) {
	amount := 6
	secret := strings.Repeat("a", 32)

	gen, err := token.NewGenerator(
		option.OptionWithSecret(secret),
		option.OptionWithPeriod(time.Second),
		option.OptionWithAmount(amount),
		option.OptionWithEmojies(emojies.CATS),
		option.OptionWithTime(time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)),
	)
	if err != nil {
		t.Error("Expected to not return an error when creating new generator:", err)
	}

	code, err := gen.Next()
	if err != nil {
		t.Error("Expected to not return an error when creating new code:", err)
	}

	errStr := "Expected %s to be %s not %s"
	correct := "🙀😾😹🙀😼😹"
	if !gen.Validate(correct) {
		t.Errorf(errStr, "code", correct, code)
	}

	correct = "187603"
	if !gen.ValidateToken(correct) {
		t.Errorf(errStr, "token", correct, code.Token())
	}
}
