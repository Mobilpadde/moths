package emojies

import "github.com/enescakir/emoji"

var (
	DOGS_SLICE = []string{
		emoji.DogFace.String(),
		emoji.Dog.String(),
		emoji.GuideDog.String(),
		// emoji.ServiceDog.String(), // Uses multiple "emojies" in one
		emoji.Poodle.String(),
		emoji.Wolf.String(),
		emoji.HotDog.String(), // Wait, what?
	}

	DOGS     = ToEmojies(DOGS_SLICE)
	DOGS_CAT = ToEmojies(append(
		DOGS_SLICE,
		emoji.KissingCat.String(),
	))
)
