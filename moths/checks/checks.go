package checks

import (
	"time"

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

func CheckInterval(interval time.Duration) error {
	if interval.Milliseconds() < 1000 {
		return errs.ErrInterval1
	}

	return nil
}
