package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// type problem struct {
// 	question string
// 	answer   string
// }

func GetUserInputAndScan(filename string) {

	body, err := os.Open(filename)
	if err != nil {
		fmt.Printf("could not open file: %v", err)
		return
	}
	defer body.Close()
	reader := csv.NewReader(body)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("could not read line: %v", err)
			return
		}

		fmt.Println(line)
	}
	fmt.Println("we good!!!!")
}

// func parseCSVInput(csvInput string) {

// }
