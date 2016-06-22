package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	httpPort 	= flag.String("port", ":8080", "Listen address")
	serverEnv 	= flag.Bool("isDev", true, "Server environment mode")
)

func main() {
	flag.Parse()

	router := NewRouter()
	router.PathPrefix("/media/").Handler(http.StripPrefix("/media/", http.FileServer(http.Dir("./media/"))))

	log.Fatal(http.ListenAndServe(*httpPort, router))
}
