package emojies

import "github.com/enescakir/emoji"

var CATS = Emojies{
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

var CATS_DOG = append(
	CATS,
	emoji.DogFace.String(),
)
