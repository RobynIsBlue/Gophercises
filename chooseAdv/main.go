package main

import (
	"log"
	"os"
	"text/template"
)

type story struct {
	Intro chapter `json:"intro"`
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
	story, err := createDecodedMap("story.json")
	if err != nil {
		log.Fatalf("could not find story: %v", err)
	}
	tmpl, err := template.New("test").Parse("{{.Intro}} IS THE CHAPTER AND {{.Intro.Story}} IS THE STORY")
	tmpl.Execute(os.Stdout, story)
}
