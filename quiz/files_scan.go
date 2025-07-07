package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func GetUserInputAndScan(filename string, timePerQuestion int) {
	body, err := os.Open(filename)
	if err != nil {
		fmt.Printf("could not open file: %v", err)
		return
	}
	defer body.Close()
	reader := csv.NewReader(body)
	readerInput := bufio.NewReader(os.Stdin)
	numTotal := 0
	numCorrect := 0

	questionAnsweredBool := make(chan bool)
	for {

		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("could not read line: %v", err)
			return
		}
		numTotal++

		lineProblem := parseCSVInput(line)
		fmt.Printf("Problem %d: %v   ", numTotal, lineProblem.question)
		input, err := readerInput.ReadString('\n')
		input = strings.TrimSpace(input)

		if err != nil {
			fmt.Printf("could not read input: %v", err)
			return
		}
		if string(input) == lineProblem.answer {
			numCorrect++
		}
	}
	fmt.Printf("--------------------\nCorrect: %d\nTotal: %d\nPercentage: %.2f%%", numCorrect, numTotal, (float32(numCorrect)/float32(numTotal))*100)
}

func parseCSVInput(csvInput []string) problem {
	return problem{
		question: strings.Join(csvInput[:len(csvInput)-1], ""),
		answer:   csvInput[len(csvInput)-1],
	}
}
