package main

import (
	"encoding/csv"
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

	r := csv.NewReader(file)

	lines, err := r.ReadAll()

	problems := paresLines(lines)

	correct := 0

	for i, n := range problems {
		fmt.Printf("QUESTION #%d: %s =\n", i+1, n.q)
		var ans string
		fmt.Scanln(&ans)

		if ans == n.a {
			correct++
		}
	}

	fmt.Printf("%d out of %d is Correct \n", correct, len(problems))
	if correct == len(problems) {
		fmt.Println("All Correct, Great Job!")
	} else {
		fmt.Println("Try Again!")
	}

}

type problem struct {
	q string
	a string
}

func paresLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, p := range lines {
		ret[i] = problem{
			q: p[0],
			a: p[1],
		}
	}
	return ret
}

func exit(errMessage string) {
	fmt.Println(errMessage)
	os.Exit(1)
}
