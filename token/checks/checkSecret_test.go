package checks

import (
	"strings"
	"testing"
)

func TestCheckSecret(t *testing.T) {
	secret := strings.Repeat("a", 32)
	if err := CheckSecret(secret); err != nil {
		t.Error("Expected to not return an error when checking ok. secret:", err)
	}
}

func TestCheckSecretError(t *testing.T) {
	secret := "0"
	if err := CheckSecret(secret); err != nil {
		t.Error("Expected to return an error when checking not ok. secret:", err)
	}
}
