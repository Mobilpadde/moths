package emojies

import "github.com/enescakir/emoji"

var (
	CATS_SLICE = []string{
		emoji.GrinningCat.String(),
		emoji.GrinningCatWithSmilingEyes.String(),
		emoji.CatWithTearsOfJoy.String(),
		emoji.SmilingCatWithHeartEyes.String(),
		emoji.CatWithWrySmile.String(),
		emoji.KissingCat.String(),
		emoji.WearyCat.String(),
		emoji.CryingCat.String(),
		emoji.PoutingCat.String(),
	}

	CATS     = ToEmojies(CATS_SLICE)
	CATS_DOG = ToEmojies(
		append(
			CATS_SLICE,
			emoji.DogFace.String(),
		),
	)
)
