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
	fileName := flag.String("f", "problems.csv", "redefines the csv file used for problems")
	flag.Parse()
	fmt.Println(http.DetectContentType([]byte(*fileName)))
	if filepath.Ext(*fileName) != ".csv" {
		panic("file is not csv")
	}
	GetUserInputAndScan(*fileName)
}
