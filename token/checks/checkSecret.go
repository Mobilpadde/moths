package checks

import "github.com/Mobilpadde/moths/v4/token/errs"

func CheckSecret(secret string) error {
	if len(secret) == 0 {
		return errs.ErrSecretLength
	}

	return nil
}
