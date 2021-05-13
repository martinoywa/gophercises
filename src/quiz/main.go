package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type questionAnswer struct {
	Question string
	Answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func getScore(csvLines [][]string, timer time.Timer) {
	var userAnswer string
	var scoreCount int
	problemCount := len(csvLines)

	for i, line := range csvLines {
		qa := questionAnswer{
			Question: line[0],
			Answer:   strings.TrimSpace(line[1]), // helps remove spaces if exists e.g 5+5, 10
		}

		select {
		case <-timer.C:
			fmt.Printf("You scored %d / %d\n", scoreCount, problemCount)
			return
		default:
			fmt.Printf("Problem #%d: %s = ", i+1, qa.Question)
			fmt.Scan(&userAnswer)

			if userAnswer == qa.Answer {
				scoreCount += 1
			}
		}
	}
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file with question answer format")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	csvFile, err := os.Open(*csvFileName)
	defer csvFile.Close()
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", err))
	} else {
		fmt.Println("Success reading CSV file.")
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		exit(fmt.Sprintln(err))
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	getScore(csvLines, *timer)
	
}
