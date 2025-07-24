package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type story struct {
	Story map[string]chapter
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type chapter struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type world struct {
	World string
}

func main() {
	base, _ := template.ParseFiles("base.html")
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		base.Execute(w, world{World: "worldy"})
	})
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css"))))
	fmt.Println("serving on port http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

