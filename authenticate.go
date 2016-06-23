package main

import (
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/sfreiberg/gotwilio"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func generatePin() string {
	pin := ""
	i := 0

	for ; i < 4; i++ {
		pin = pin + strconv.Itoa(rand.Intn(10))
	}
	println("pin:", pin)
	return pin
}

func fixPhone(phone string) string {
	reg, err := regexp.Compile("[^0-9+]")
	if (err != nil) {
		log.Fatal(err)
	}

	number := reg.ReplaceAllString(phone, "");
	regNA, err := regexp.MatchString("^[0-9]{10}$", number)
	if (err != nil) {
		log.Fatal(err)
	}

	regPl, err := regexp.MatchString("^+", number)
	if (err != nil) {
		log.Fatal(err)
	}

	if (regNA) {
	  number = "+1" + number;
	} else if (regPl) {
	  number = "+" + number;
	}
	println(number)
	return number
}

func sendPin(phone string) {
	twilio := gotwilio.NewTwilioClient(*twilioSid, *twilioToken)
	from := *twilioPhone
	to := fixPhone(phone)
	message := "Your PIN: " + generatePin()
	twilio.SendSMS(from, to, message, "", "")
}
