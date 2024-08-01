package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {

	problemsFile := flag.String("file", "problems.csv", "file name which has all the problems")
	limit := flag.Int("limit", 10, "time to complete the quiz")

	flag.Parse()

	problemFile, err := os.Open(*problemsFile)

	if err != nil {
		fmt.Println("Cannot open the file.")
		os.Exit(1)
	}

	r := csv.NewReader(problemFile)
	lines, err := r.ReadAll()

	if err != nil {
		fmt.Println("CSV file parsing error")
	}

	problems := parseCsvLines(lines)

	timer1 := time.NewTimer(time.Duration(*limit) * time.Second)

	correct := 0

problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d : %s = ", i+1, p.q)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer1.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("You have scored %d out of %d problems!", correct, len(problems))
}

func parseCsvLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, p := range lines {
		ret[i] = problem{p[0], strings.TrimSpace(p[1])}
	}

	return shuffle(ret)
}

func shuffle(problems []problem) []problem {
	rand.Shuffle(len(problems), func(i, j int) {
		problems[i], problems[j] = problems[j], problems[i]
	})
	return problems
}
