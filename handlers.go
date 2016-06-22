package main

import (
	"net/http"

	"github.com/unrolled/render"
)

var r = render.New(render.Options{
	Directory: "templates",
	Extensions: []string{".html"},
	IsDevelopment: *serverEnv,
})

func Index(w http.ResponseWriter, req *http.Request) {
	r.HTML(w, http.StatusOK, "index", nil)
}

func Directory(w http.ResponseWriter, req *http.Request) {
	r.HTML(w, http.StatusOK, "directory", nil)
}
