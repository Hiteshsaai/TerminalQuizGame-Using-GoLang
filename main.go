package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "The csv file with a formate of question,answer")
	flag.Parse()

	file, err := os.Open(*csvFile)

	if err != nil {
		message := fmt.Sprintf("Filed to open the CSV File %s \n", *csvFile)
		exit(message)
	}

	_ = file

}

func exit(errMessage string) {
	fmt.Println(errMessage)
	os.Exit(1)
}
