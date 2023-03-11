package main

import (
	"log"
	"os"
	"time"

	"github.com/Mobilpadde/moths/v5/token"
	"github.com/Mobilpadde/moths/v5/token/emojies"
	"github.com/Mobilpadde/moths/v5/token/option"
)

func main() {
	secret := os.Getenv("MOTHS_SECRET")
	amount := 8

	validationInterval := time.Second * 3
	generationInterval := time.Second * 4

	generationTicker := time.NewTicker(generationInterval)
	validationTicker := time.NewTicker(validationInterval)

	var err error
	var gen *token.Generator
	if gen, err = token.NewGenerator(
		option.OptionWithSecret(secret),
		option.OptionWithPeriod(generationInterval),
		option.OptionWithAmount(amount),
		option.OptionWithEmojies(emojies.CATS),
		option.OptionWithTime(time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)),
	); err != nil {
		log.Fatalln(err)
	}

	log.Printf("Setup with validation complete")
	log.Printf(
		"A new code is generated every %s, and validation happens %s after generation",
		generationInterval,
		validationInterval,
	)
	log.Printf("Every code is %d emoji long", amount)

	for {
		log.Println()
		code, err := gen.Next()
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf(`Your code is "%s" and code is %s`, code.SpacedString(), code.Token())
		<-validationTicker.C

		log.Printf("Is this still valid after %s? %t", validationInterval, gen.Validate(code.String()))
		<-validationTicker.C

		log.Printf("Is this still valid after %s? %t", validationInterval*2, gen.Validate(code.String()))
		<-generationTicker.C
	}
}
