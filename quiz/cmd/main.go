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
	var nFlag = flag.Int("n", 10, "Number of seconds to wait for user input")
	var fName = flag.String("f", DefaultQuizFile, "Absolute path of the quiz file")
	flag.Parse()
	err := quiz.Quiz(*fName, time.Duration(*nFlag)*time.Second)
	if err != nil {
		fmt.Printf("Quiz terminated with error %s", err)
	}
}
