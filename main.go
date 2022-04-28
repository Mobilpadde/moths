package main

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/enescakir/emoji"
)

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)

	data, err := ioutil.ReadFile("secret.pem")
	if err != nil {
		log.Fatalln(err)
	}

	secret := string(data)
	interval := time.Now().Unix() / 5

	otp := GetHOTPToken(secret, interval)
	moth, err := MothsGenerator(
		otp,
		5,
		emoji.GrinningCat.String(),
		emoji.GrinningCatWithSmilingEyes.String(),
		emoji.CatWithTearsOfJoy.String(),
		emoji.SmilingCatWithHeartEyes.String(),
		emoji.CatWithWrySmile.String(),
		emoji.KissingCat.String(),
		emoji.WearyCat.String(),
		emoji.CryingCat.String(),
		emoji.PoutingCat.String(),
	)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("moth-otp: %s => %s\n", moth, otp)
}
