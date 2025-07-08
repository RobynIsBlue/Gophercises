package main

import (
	// "encoding/csv"
	"flag"
	"fmt"
	"net/http"
	"path/filepath"
)

// type fileParams struct {
// 	fileName string
// }

func main() {
	timeTotal := flag.Int("o", 30, "choose time for whole quiz")
	fileName := flag.String("f", "problems.csv", "redefines the csv file used for problems")
	timePerQuestion := flag.Int("t", 4, "choose time per question")
	flag.Parse()
	fmt.Println(http.DetectContentType([]byte(*fileName)))
	if filepath.Ext(*fileName) != ".csv" {
		panic("file is not csv")
	}
	GetUserInputAndScan(*fileName, *timePerQuestion, *timeTotal)
}
