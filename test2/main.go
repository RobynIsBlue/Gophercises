package main

import (
	"net/http"

	"github.com/a-h/templ"
)

// howdy pardnur !!!!!!!!!!!!!1

func main() {
	comp := hello("world")
	http.FileServer(http.Dir("/static/styles/"))
	http.Handle("/", templ.Handler(comp))
	http.ListenAndServe(":8080", nil)
}
