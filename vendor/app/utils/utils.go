package utils

import (
	"flag"

	"github.com/gorilla/sessions"
	"github.com/unrolled/render"
)

type Flag struct {
	HttpPort		string
	ServerEnv		bool
	TwilioSid		string
	TwilioToken		string
	TwilioPhone		string
	CookieSecret	string
}

var (
	httpPort		= flag.String("port", ":8080", "Listen address")
	serverEnv		= flag.Bool("isDev", true, "Server environment mode")
	twilioSid		= flag.String("twilioSid", "111", "Twilio SID")
	twilioToken		= flag.String("twilioToken", "111", "Twilio token")
	twilioPhone		= flag.String("twilioPhone", "+15555555", "Twilio phone number")
	cookieSecret	= flag.String("cookie", "secret", "Session cookie secret")
	store			= sessions.NewCookieStore([]byte(*cookieSecret))

	r				= render.New(render.Options{
						Directory: "templates",
						Extensions: []string{".tmpl"},
						Layout: "layout",
						IsDevelopment: *serverEnv,
					})
)

func init() {
	flag.Parse()
}

func GetRender() *render.Render {
	return r
}

func GetSession() *sessions.CookieStore {
	return store
}

func GetFlags() *Flag {
	flags := &Flag{HttpPort: *httpPort,
		ServerEnv: *serverEnv,
		TwilioSid: *twilioSid,
		TwilioToken: *twilioToken,
		TwilioPhone: *twilioPhone,
		CookieSecret: *cookieSecret}
	return flags
}