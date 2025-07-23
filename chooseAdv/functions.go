package main

import (
	"encoding/json"
	"fmt"
	"io"
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

func readStory(path string) (story, error) {
	var storyThing map[string]interface{}
	body, err := os.Open(path)
	if err != nil {
		return story{}, err
	}
	defer body.Close()
	byteValue, _ := io.ReadAll(body)
	json.Unmarshal(byteValue, &storyThing)
	fmt.Println(storyThing["intro"])
	return story{}, nil
}