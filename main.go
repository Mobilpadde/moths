package main

import (
	"log"
	"os"
	"time"

	"moths/moths"
	"moths/moths/emojies"
)

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)

	secret := os.Getenv("MOTHS_SECRET")
	interval := 10
	amount := 10

	validInterval := time.Second * 7

	var err error
	var gen *moths.Moths
	if gen, err = moths.NewMoths(
		moths.WithSecret(secret),
		moths.WithInterval(interval),
		moths.WithAmount(amount),
		moths.WithEmojies(emojies.DOGS),
	); err != nil {
		log.Fatalln(err)
	}
	ticker := time.NewTicker(time.Second * time.Duration(interval))
	for {
		otp, err := gen.Next()
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf(`Your moth is "%s" and code is %s`, otp.SpacedString(), otp.Token())

		<-time.NewTicker(validInterval).C
		log.Printf("Is this still valid after %s? %t", validInterval, gen.Validate(otp.Token()))

		<-ticker.C
	}
}
