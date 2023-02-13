package main

import (
	"flag"
	"fmt"
	"quiz/quiz"
	"time"
)

var (
	DefaultQuizFile = "C:\\Users\\BaskarHome\\Coding\\golang\\golang-examples\\quiz\\cmd\\problem.csv"
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
