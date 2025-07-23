package main

import (
	"fmt"
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

func main() {
	// _, err := readStory("story.json")
	// if err != nil {
	// 	log.Fatalf("could not find story: %v", err)
	// }
	// comp := hello("world")
	// comp.Render(context.Background(), os.Stdout)

	// http.Handle("/", templ.Handler(comp))
	fmt.Println("serving on port http://localhost:8080")
	http.ListenAndServe(":8080", )
}

