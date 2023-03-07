package checks

import "github.com/Mobilpadde/moths/v5/token/errs"

func CheckSecret(secret string) error {
	if len(secret) == 0 {
		return errs.ErrSecretLength
	}

	return nil
}
