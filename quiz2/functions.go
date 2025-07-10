package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func filePathToReader(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("could not open filepath: %v", err)
	}
	csvReader := csv.NewReader(file)
	contents, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("could not read contents: %v", err)
	}
	return contents
}

func quiz(problems [][]string, timeOverall int) {
	var correct int
	timerOverall := time.NewTimer(time.Duration(timeOverall) * time.Second)
	inputChan := make(chan string)
Done:
	for i, problem := range problems {
		fmt.Printf("Problem %d: %v\n", i+1, strings.Join(problem[:len(problem)-1], ""))
		go getInput(inputChan)
		select {
		case <-timerOverall.C:
			break Done
		case inputR := <-inputChan:
			if inputR == problem[len(problem)-1] {
				correct++
			}
		}
	}
	fmt.Printf("\n-----------------\nCorrect: %v\nTotal: %v\nPercent: %.2f%%",
		correct,
		len(problems),
		float64(correct)/(float64(len(problems)))*100)
}

func getInput(inputChan chan string) {
	var input string
	fmt.Scanln(&input)
	inputChan <- input
}
