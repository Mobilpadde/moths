package emojies

import "github.com/enescakir/emoji"

var dogs = []string{
	emoji.DogFace.String(),
	emoji.Dog.String(),
	emoji.GuideDog.String(),
	// emoji.ServiceDog.String(), // Uses multiple "emojies" in one
	emoji.Poodle.String(),
	emoji.Wolf.String(),
	emoji.HotDog.String(), // Wait, what?
}

var DOGS = ToEmojies(dogs)

var DOGS_CAT = ToEmojies(append(
	dogs,
	emoji.KissingCat.String(),
))
