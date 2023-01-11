package emojies

import "github.com/enescakir/emoji"

var DOGS = Emojies{
	emoji.DogFace.String(),
	emoji.Dog.String(),
	emoji.GuideDog.String(),
	emoji.ServiceDog.String(),
	emoji.Poodle.String(),
	emoji.Wolf.String(),
	emoji.HotDog.String(), // Wait, what?
}

var DOGS_CAT = append(
	DOGS,
	emoji.KissingCat.String(),
)
