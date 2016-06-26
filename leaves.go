package main

import (
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/unrolled/secure"
	"github.com/urfave/negroni"

	"app/routes"
	"app/utils"
)

func main() {
	flags := utils.GetFlags()
	router := routes.NewRouter()
	router.PathPrefix("/media/").Handler(http.StripPrefix("/media/", http.FileServer(http.Dir("./media/"))))

	csrf := csrf.Protect(
		[]byte(utils.GetFlags().CsrfSecret),
		csrf.Secure(!flags.IsDev),
	)

	csp := secure.New(secure.Options{
		AllowedHosts: []string{"fonts.googleapis.com"},
		FrameDeny: true,
		IsDevelopment: flags.IsDev,
	})

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.HandlerFunc(csp.HandlerFuncWithNext))
	n.UseHandler(router)

	log.Fatal(http.ListenAndServe(flags.HttpPort, csrf(n)))
}
