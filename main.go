package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var (
	httpPort 		= flag.String("port", ":8080", "Listen address")
	serverEnv 		= flag.Bool("isDev", true, "Server environment mode")
	twilioSid 		= flag.String("twilioSid", "111", "Twilio SID")
	twilioToken 	= flag.String("twilioToken", "111", "Twilio token")
	twilioPhone		= flag.String("twilioPhone", "+15555555", "Twilio phone number")
	cookieSecret 	= flag.String("cookie", "secret", "Session cookie secret")
	store 			= sessions.NewCookieStore([]byte(*cookieSecret))

	r				= render.New(render.Options{
						Directory: "templates",
						Extensions: []string{".tmpl"},
						Layout: "layout",
						IsDevelopment: *serverEnv,
					})
)

func main() {
	flag.Parse()

	router := NewRouter()
	router.PathPrefix("/media/").Handler(http.StripPrefix("/media/", http.FileServer(http.Dir("./media/"))))

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	log.Fatal(http.ListenAndServe(*httpPort, n))
}

func Index(w http.ResponseWriter, req *http.Request) {
	session, err := store.Get(req, "phone")
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s := false

	if (session.Values["phone"] != nil) {
		fmt.Println(*session)
		s = true
	}

	r.HTML(w, http.StatusOK, "index", map[string]interface{}{
		"session": s,
	})
}

func Directory(w http.ResponseWriter, req *http.Request) {
	r.HTML(w, http.StatusOK, "directory", nil)
}

func Authenticate(w http.ResponseWriter, req *http.Request) {
	if (req.Method == "POST") {
		session, err := store.Get(req, "phone")
		if (err != nil) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		decoder := req.FormValue("phone")
		phone := sendPin(decoder)
		session.Values["phone"] = phone;
		session.Save(req, w)

		http.Redirect(w, req, "/validate", 301)
	} else {
		r.HTML(w, http.StatusOK, "authenticate", nil)
	}
}

func Validate(w http.ResponseWriter, req *http.Request) {
	if (req.Method == "POST") {
		session, err := store.Get(req, "phone")
		if (err != nil) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		pin := req.FormValue("pin")
		phone := session.Values["phone"].(string)
		pinVerify := validatePin(pin, phone)

		if (pinVerify) {
			http.Redirect(w, req, "/", 301)
		} else {
			r.HTML(w, http.StatusOK, "validate", nil)
		}
	} else {
		r.HTML(w, http.StatusOK, "validate", nil)
	}
}
