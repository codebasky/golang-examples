package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/codebasky/golang-examples/quiz/quiz"
)

var (
	DefaultQuizFile = "problem.csv"
)

func main() {
	var waitTime = flag.Int("w", 10, "Number of seconds to wait for user input")
	var fileName = flag.String("f", DefaultQuizFile, "Absolute path of the quiz file")
	flag.Parse()
	err := quiz.Quiz(*fileName, time.Duration(*waitTime)*time.Second)
	if err != nil {
		fmt.Printf("Quiz terminated with error %s", err)
	}
}
