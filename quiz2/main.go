package main

import "flag"

func main() {
	filePath := flag.String("p", "problems.csv", "sets a path to a csv file containing problems")
	time := flag.Int("t", 2, "sets the time for the file")
	flag.Parse()
	quiz(filePathToReader(*filePath), *time)
}
