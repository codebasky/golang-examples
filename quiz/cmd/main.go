package main

import (
	"fmt"
	"quiz/quiz"
)

func main() {
	err := quiz.Quiz("problem.csv")
	if err != nil {
		fmt.Printf("Quiz terminated with error %s", err)
	}
}
