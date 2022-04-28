package moths

import "github.com/enescakir/emoji"

var (
	ALPHABET_CATSDOG = []string{
		emoji.GrinningCat.String(),
		emoji.GrinningCatWithSmilingEyes.String(),
		emoji.CatWithTearsOfJoy.String(),
		emoji.SmilingCatWithHeartEyes.String(),
		emoji.CatWithWrySmile.String(),
		emoji.KissingCat.String(),
		emoji.WearyCat.String(),
		emoji.CryingCat.String(),
		emoji.PoutingCat.String(),
		emoji.DogFace.String(),
	}
)

type Moths struct {
	secret   string
	interval int
	alphabet []string
}
