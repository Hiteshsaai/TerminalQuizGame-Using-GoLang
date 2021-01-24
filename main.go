package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "The csv file with a formate of question,answer")
	timeLimiter := flag.Int("limit", 30, "The timer runs in seconds")
	flag.Parse()

	file, err := os.Open(*csvFile)

	Timer := time.NewTimer(time.Duration(*timeLimiter) * time.Second)

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
		var ansCha = make(chan string)
		go func() {
			fmt.Scanln(&ans)
			ansCha <- ans
		}()
		select {
		case <-Timer.C:
			fmt.Println("Time Up!")
			fmt.Printf("%d out of %d is Correct \n", correct, len(problems))
			return
		case answer := <-ansCha:
			if answer == n.a {
				correct++
			}
		}

	}

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
			a: strings.TrimSpace(p[1]),
		}
	}
	return ret
}

func exit(errMessage string) {
	fmt.Println(errMessage)
	os.Exit(1)
}
