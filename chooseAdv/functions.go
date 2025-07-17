package main

import (
	"encoding/json"
	"os"
)

func createDecodedMap(path string) (story, error) {
	var decodedStory story
	body, err := os.Open(path)
	if err != nil {
		return story{}, err
	}
	defer body.Close()
	err = json.NewDecoder(body).Decode(&decodedStory)
	if err != nil {
		return story{}, err
	}
	return decodedStory, nil
}
