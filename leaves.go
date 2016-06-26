package main

import (
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/urfave/negroni"

	"app/routes"
	"app/utils"
)

func main() {
	flags := utils.GetFlags()
	router := routes.NewRouter()
	router.PathPrefix("/media/").Handler(http.StripPrefix("/media/", http.FileServer(http.Dir("./media/"))))

	CSRF := csrf.Protect(
		[]byte(utils.GetFlags().CsrfSecret),
		csrf.Secure(!utils.GetFlags().IsDev),
	)

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	log.Fatal(http.ListenAndServe(flags.HttpPort, CSRF(n)))
}
