package main

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/nu7hatch/gouuid"
	"github.com/patrickmn/go-cache"
	"github.com/sfreiberg/gotwilio"
)

// 5 minute cache for pin
var c = cache.New(5*time.Minute, 30*time.Second)

type Profile struct {
	Uid			*uuid.UUID
	Name		string
	Phone		string
}

type ProfileList []*Profile

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

func sendPin(phone string) string {
	twilio := gotwilio.NewTwilioClient(*twilioSid, *twilioToken)
	from := *twilioPhone
	to := fixPhone(phone)
	pin := generatePin()
	message := "Your PIN: " + pin
	twilio.SendSMS(from, to, message, "", "")

	c.Set(to, pin, cache.DefaultExpiration)
	return to
}

func validatePin(pin string, phone string) bool {
	phonePin, found := c.Get(phone)
	if found {
		p := phonePin.(string)
		if (p == pin) {
			return true
		}
	}
	return false
}

func createProfile(phone string) {
	u, err := uuid.NewV4()
	if (err != nil) {
		log.Fatal(err)
	}

	var list ProfileList
	list = append(list, &Profile{Uid: u, Name: "???", Phone: phone})
	for _, e := range list {
        fmt.Println("  ", *e)
    }
}
