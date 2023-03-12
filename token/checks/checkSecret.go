package checks

import "github.com/Mobilpadde/moths/v6/token/errs"

// CheckSecret is used to check
// if the provided secret is "strong" enough.
func CheckSecret(secret string) error {
	if len(secret) == 0 {
		return errs.ErrSecretLength
	}

	return nil
}
