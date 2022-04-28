package main

import (
	"log"
	"moths/moths"
	"os"
	"strings"
	"time"
)

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)

	go func() {
		interval := 5

		secret := os.Getenv("MOTHS_SECRET")
		gen, err := moths.NewMoths(secret, interval)
		if err != nil {
			log.Fatalln(err)
		}

		ticker := time.NewTicker(time.Second * time.Duration(interval))
		for {
			select {
			case <-ticker.C:
				moth, err := gen.Next(6, moths.ALPHABET_CATSDOG)
				if err != nil {
					log.Fatalln(err)
				}

				joint := strings.Join(strings.Split(moth, ""), " ")
				log.Printf("moth-otp: %s\n", joint)
			}
		}
	}()

	select {}
}
