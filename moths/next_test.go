package moths

import (
	"strings"
	"testing"
	"time"

	"github.com/Mobilpadde/moths/moths/emojies"
)

func TestNextOk(t *testing.T) {
	amount := 6
	secret := strings.Repeat("a", 32)

	gen, err := NewMoths(
		WithSecret(secret),
		WithInterval(time.Second),
		WithAmount(amount),
		WithEmojies(emojies.CATS),
		WithTime(time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)),
	)
	if err != nil {
		t.Error("Expected to not return an error when creating new Moths:", err)
	}

	otp, err := gen.Next()
	if err != nil {
		t.Error("Expected to not return an error when creating new OTP:", err)
	}

	correct := "158415"
	if otp.Token() != correct {
		t.Error("Expected token to be", correct, "not", otp.Token())
	}

	correct = "🙀🙀😺😼😽😻"
	if otp.String() != correct {
		t.Error("Expected moth to be", correct, "not", otp.String())
	}
}
