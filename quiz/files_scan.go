package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func GetUserInputAndScan(filename string, timePerQuestion, timeTotal int) {
	body, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("could not open file: %v", err)
		return
	}
	splitBody := strings.Split(string(body), "\r\n")
	readerInput := bufio.NewReader(os.Stdin)
	var numCorrect int
	timer := time.NewTimer(time.Second * time.Duration(timePerQuestion))
	quizChan := make(chan string)
	bigTimer := time.NewTimer(time.Second * time.Duration(timeTotal))
bigLoop:
	for i, line := range splitBody {
		timer.Reset(time.Second * time.Duration(timePerQuestion))
		lineProblem := parseCSVInput(line)
		go func() {
			fmt.Printf("Problem %d: %v   ", i+1, lineProblem.question)
			input, err := readerInput.ReadString('\n')
			input = strings.TrimSpace(input)

			if err != nil {
				fmt.Printf("could not read input: %v", err)
				return
			}
			quizChan <- input
		}()
		select {
		case <-bigTimer.C:
			break bigLoop
		case <-timer.C:
			fmt.Println()
		case input := <-quizChan:
			if input == lineProblem.answer {
				numCorrect++
			}
		}
	}
	fmt.Printf("--------------------\nCorrect: %d\nTotal: %d\nPercentage: %.2f%%", numCorrect, numTotal, (float32(numCorrect)/float32(numTotal))*100)
}

func parseCSVInput(csvInput string) problem {
	reader := csv.NewReader(strings.NewReader(csvInput))
	readerInput, err := reader.Read()
	if err != nil {
		fmt.Printf("could not parse csv input: %v\n", err)
		return problem{}
	}
	return problem{
		question: strings.Join(readerInput[:len(readerInput)-1], ""),
		answer:   readerInput[len(readerInput)-1],
	}
}