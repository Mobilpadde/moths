package checks

import (
	"github.com/Mobilpadde/moths/v5/token/emojies"
	"github.com/Mobilpadde/moths/v5/token/errs"
)

func CheckEmojies(emojies emojies.Emojies) error {
	if len(emojies) < 1 {
		return errs.ErrEmojies1
	}

	return nil
}
