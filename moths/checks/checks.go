package checks

import (
	"time"

	"github.com/Mobilpadde/moths/moths/emojies"
	"github.com/Mobilpadde/moths/moths/errs"
)

func CheckSecret(secret string) error {
	if len(secret) == 0 {
		return errs.ErrSecretLength
	}

	return nil
}

func CheckSecretKey(key []byte) error {
	// if len(key) != 20 {
	// 	return errs.ErrKey20
	// }

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

func CheckAmount(amount int) error {
	if amount < 1 {
		return errs.ErrAmount1
	}

	return nil
}
