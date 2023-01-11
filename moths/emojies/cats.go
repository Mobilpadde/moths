package emojies

import "github.com/enescakir/emoji"

var cats = []string{
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

var CATS = ToEmojies(cats)

var CATS_DOG = ToEmojies(append(
	cats,
	emoji.DogFace.String(),
))
