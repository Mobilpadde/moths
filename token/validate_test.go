package token_test

import (
	"log"
	"strings"
	"time"

	"github.com/Mobilpadde/moths/v5/token"
	"github.com/Mobilpadde/moths/v5/token/emojies"
	"github.com/Mobilpadde/moths/v5/token/option"
)

func ExampleGenerator_Validate() {
	amount := 6
	secret := strings.Repeat("a", 32)

	var err error
	var gen *token.Generator
	if gen, err = token.NewGenerator(
		option.OptionWithSecret(secret),
		option.OptionWithPeriod(time.Second),
		option.OptionWithAmount(amount),
		option.OptionWithEmojies(emojies.CATS),
	); err != nil {
		log.Fatalln(err)
	}

	code, _ := gen.Next()
	gen.Validate(code.String()) // This is true, as it's validated within specified period.
}
