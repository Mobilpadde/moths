package main

import (
	"log"
	"os"
	"time"

	"github.com/Mobilpadde/moths/v6/token"
	"github.com/Mobilpadde/moths/v6/token/emojies"
	"github.com/Mobilpadde/moths/v6/token/option"
)

func main() {
	secret := os.Getenv("MOTHS_SECRET") // created as specified in the README: https://github.com/Mobilpadde/moths#how-
	amount := 8                         // we want to generate a code with 8 emojies

	validationInterval := time.Second * 3
	generationInterval := time.Second * 4

	// Instantiate a new generator
	gen, err := token.NewGenerator(
		option.OptionWithSecret(secret),
		option.OptionWithPeriod(generationInterval),
		option.OptionWithAmount(amount),
		option.OptionWithEmojies(emojies.CATS),
		option.OptionWithTime(time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)),
	)
	if err != nil {
		log.Fatalln(err)
	}

	// Check if everything is working
	if err := gen.Check(); err != nil {
		log.Fatalln(err)
	}

	// Exports the generator's options
	s := gen.Export()

	// Instantiate a second generator
	gen2, _ := token.NewGenerator()

	// Import the encoded options into the new generator
	if err = gen2.Import(s); err != nil {
		log.Println(err)
	}

	// Check if everything is still working
	if err := gen2.Check(); err != nil {
		log.Fatalln(err)
	}

	log.Printf("Setup with validation complete")
	log.Printf(
		"A new code is generated every %s, and validation happens %s after generation",
		generationInterval,
		validationInterval,
	)
	log.Printf("Every code is %d emoji long", amount)

	// Flow:
	//
	// 10 Validate a code after 3 seconds (validationInterval)
	// 20 Try validating the code again after 3 more seconds (validationInterval)
	// 30 Generate a new code
	//
	// 40 GOTO 10
	for {
		log.Println()
		code, err := gen2.Next()
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf(`Your code is "%s" and code is %s`, code.SpacedString(), code.Token())
		<-time.NewTicker(validationInterval).C

		log.Printf("Is this still valid after %s? %t", validationInterval, gen2.Validate(code.String()))
		<-time.NewTicker(validationInterval).C

		log.Printf("Is this still valid after %s? %t", validationInterval*2, gen2.Validate(code.String()))
	}
}
