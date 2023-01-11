package checks

import (
	"moths/moths/emojies"
	"moths/moths/errs"
)

func CheckSecret(secret string) error {
	if len(secret) != 32 {
		return errs.ErrSecret32
	}

	return nil
}

func CheckSecretKey(key []byte) error {
	if len(key) != 20 {
		return errs.ErrKey32
	}

	return nil
}

func CheckEmojies(emojies emojies.Emojies) error {
	if len(emojies) < 1 {
		return errs.ErrEmojies1
	}

	return nil
}

func CheckInterval(interval int) error {
	if interval < 1 {
		return errs.ErrInterval1
	}

	return nil
}
