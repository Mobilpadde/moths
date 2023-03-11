package checks_test

import (
	"strings"
	"testing"

	"github.com/Mobilpadde/moths/v5/token/checks"
)

func TestCheckSecret(t *testing.T) {
	secret := strings.Repeat("a", 32)
	if err := checks.CheckSecret(secret); err != nil {
		t.Error("Expected to not return an error when checking ok. secret:", err)
	}
}

func TestCheckSecretError(t *testing.T) {
	secret := "0"
	if err := checks.CheckSecret(secret); err != nil {
		t.Error("Expected to return an error when checking not ok. secret:", err)
	}
}
