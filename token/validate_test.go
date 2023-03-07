package token

import (
	"log"
	"strings"
	"time"

	"github.com/Mobilpadde/moths/v5/token/emojies"
)

func ExampleGenerator_Validate() {
	amount := 6
	secret := strings.Repeat("a", 32)

	var err error
	var gen *Generator
	if gen, err = NewGenerator(
		OptionWithSecret(secret),
		OptionWithPeriod(time.Second),
		OptionWithAmount(amount),
		OptionWithEmojies(emojies.CATS),
	); err != nil {
		log.Fatalln(err)
	}

	code, _ := gen.Next()
	gen.Validate(code.String()) // This is true, as it's validated within specified period.
}
