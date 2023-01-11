package main

import (
	"log"
	"os"
	"time"

	"moths/moths"
	"moths/moths/emojies"
)

func main() {
	secret := os.Getenv("MOTHS_SECRET")
	amount := 10

	generationInterval := time.Second * 10
	generationTicker := time.NewTicker(generationInterval).C

	validInterval := time.Second * 7
	validationTicker := time.NewTicker(validInterval).C

	var err error
	var gen *moths.Moths
	if gen, err = moths.NewMoths(
		moths.WithSecret(secret),
		moths.WithInterval(generationInterval),
		moths.WithAmount(amount),
		moths.WithEmojies(emojies.CATS),
	); err != nil {
		log.Fatalln(err)
	}

	log.Printf("Setup with validation complete")
	log.Printf(
		"A new moth is generated every %s, and validation happens %s after generation",
		generationInterval,
		validInterval,
	)
	log.Printf("Every moth is %d emoji long", amount)
	log.Println()

	for {
		otp, err := gen.Next()
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf(`Your moth is "%s" and code is %s`, otp.SpacedString(), otp.Token())
		<-validationTicker

		log.Printf("Is this still valid after %s? %t", validInterval, gen.Validate(otp.Token()))
		log.Println()
		<-generationTicker
	}
}
