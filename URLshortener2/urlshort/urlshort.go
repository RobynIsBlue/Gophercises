package urlshort

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	fmt.Println("map handler")
	fmt.Println(pathsToUrls)
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("wah %v\n", r.URL.Path)
		mapLink, ok := pathsToUrls[r.URL.Path]
		if ok {
			fmt.Println("redirection")
			http.Redirect(w, r, mapLink, http.StatusMovedPermanently)
		} else {
			fmt.Println("fallback")
			http.HandleFunc(r.URL.Path, fallback.ServeHTTP)
		}
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return func(w http.ResponseWriter, r *http.Request) {
		yamlMap := yamlParser(yml)
		MapHandler(yamlMap, fallback).ServeHTTP(w, r)
	}, nil
}

func yamlParser(yml []byte) map[string]string {
	newMap := make(map[string]string)
	if err := yaml.Unmarshal(yml, &newMap); err != nil {
		log.Printf("oops couldn't parse yaml: %v", err)
	}
	fmt.Println(newMap)
	return newMap
}
