package main

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"app/routes"
	"app/utils"
)

func main() {
	flags := utils.GetFlags()
	router := routes.NewRouter()
	router.PathPrefix("/media/").Handler(http.StripPrefix("/media/", http.FileServer(http.Dir("./media/"))))

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	log.Fatal(http.ListenAndServe(flags.HttpPort, n))
}
