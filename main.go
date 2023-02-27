package main

import (
	"log"
	"os"
	"time"

	"github.com/Mobilpadde/moths/moths"
	"github.com/Mobilpadde/moths/moths/emojies"
)

func main() {
	secret := os.Getenv("MOTHS_SECRET")
	amount := 8

	validationInterval := time.Second * 3
	generationInterval := time.Second * 4

	generationTicker := time.NewTicker(generationInterval)
	validationTicker := time.NewTicker(validationInterval)

	var err error
	var gen *moths.Moths
	if gen, err = moths.NewMoths(
		moths.OptionWithSecret(secret),
		moths.OptionWithInterval(generationInterval),
		moths.OptionWithAmount(amount),
		moths.OptionWithEmojies(emojies.CATS),
		moths.OptionWithTime(time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)),
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
		otp, err := gen.Next()
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf(`Your code is "%s" and code is %s`, otp.SpacedString(), otp.Token())
		<-validationTicker.C

		log.Printf("Is this still valid after %s? %t", validationInterval, gen.Validate(otp.String()))
		<-validationTicker.C

		log.Printf("Is this still valid after %s? %t", validationInterval*2, gen.Validate(otp.String()))
		<-generationTicker.C
	}
}
