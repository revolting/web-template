package handlers

import (
	"fmt"
	"net/http"

	"app/authenticate"
	"app/utils"
)

var r = utils.GetRender()

func Index(w http.ResponseWriter, req *http.Request) {
	session, err := utils.GetSession().Get(req, "phone")
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
		session, err := utils.GetSession().Get(req, "phone")
		if (err != nil) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		decoder := req.FormValue("phone")
		phone := authenticate.SendPin(decoder)
		session.Values["phone"] = phone;
		session.Save(req, w)

		http.Redirect(w, req, "/validate", 301)
	} else {
		r.HTML(w, http.StatusOK, "authenticate", nil)
	}
}

func Validate(w http.ResponseWriter, req *http.Request) {
	if (req.Method == "POST") {
		session, err := utils.GetSession().Get(req, "phone")
		if (err != nil) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		pin := req.FormValue("pin")
		phone := session.Values["phone"].(string)
		pinVerify := authenticate.ValidatePin(pin, phone)

		if (pinVerify) {
			authenticate.CreateProfile(phone)

			http.Redirect(w, req, "/", 301)
		} else {
			r.HTML(w, http.StatusOK, "validate", nil)
		}
	} else {
		r.HTML(w, http.StatusOK, "validate", nil)
	}
}

func Logout(w http.ResponseWriter, req *http.Request) {
	session, err := utils.GetSession().Get(req, "phone")
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["phone"] = nil
	session.Save(req, w)
	http.Redirect(w, req, "/", 301)
}
