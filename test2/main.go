package main

import (
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	comp := hello("world")
	http.FileServer(http.Dir("static/style/"))
	http.Handle("/", templ.Handler(comp))
	http.ListenAndServe(":8080", nil)
}
