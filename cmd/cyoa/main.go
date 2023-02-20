package main

import (
	"fmt"

	"github.com/codebasky/golang-examples/cyoa/cyoa"
)

func main() {
	adv, err := cyoa.ParseJson("cyoa.json")

	if err != nil {
		fmt.Println(err)
		return
	}
	err = cyoa.StartCyoa(adv)
	if err != nil {
		fmt.Println(err)
		return
	}
}
