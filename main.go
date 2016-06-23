package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/unrolled/render"
)

var (
	httpPort 	= flag.String("port", ":8080", "Listen address")
	serverEnv 	= flag.Bool("isDev", true, "Server environment mode")
	r = render.New(render.Options{
		Directory: "templates",
		Extensions: []string{".html"},
		IsDevelopment: *serverEnv,
	})
)

func main() {
	flag.Parse()

	router := NewRouter()
	router.PathPrefix("/media/").Handler(http.StripPrefix("/media/", http.FileServer(http.Dir("./media/"))))

	log.Fatal(http.ListenAndServe(*httpPort, router))
}

func Index(w http.ResponseWriter, req *http.Request) {
	r.HTML(w, http.StatusOK, "index", nil)
}

func Directory(w http.ResponseWriter, req *http.Request) {
	r.HTML(w, http.StatusOK, "directory", nil)
}

func Authenticate(w http.ResponseWriter, req *http.Request) {
	r.HTML(w, http.StatusOK, "authenticate", nil)
}

func Validate(w http.ResponseWriter, req *http.Request) {
	r.HTML(w, http.StatusOK, "validate", nil)
}
