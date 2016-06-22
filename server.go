package main

import (
	"flag"
	"net/http"

	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

// Command-line flags.
var (
	httpPort = flag.String("port", ":8080", "Listen address")
	serverEnv = flag.Bool("isDev", true, "Server environment mode")
)

func main() {
	flag.Parse()

	r := render.New(render.Options{
		Directory: "templates",
		Extensions: []string{".html"},
		IsDevelopment: *serverEnv,
	})

	mux := http.NewServeMux()

	mux.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("media"))))

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		r.HTML(w, http.StatusOK, "index", nil)
	})

	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(*httpPort, n)
}
