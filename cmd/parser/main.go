package main

import (
	"fmt"
	"os"

	"github.com/codebasky/golang-examples/linkparser"
)

func main() {
	fNames := []string{"ex1.html", "ex2.html", "ex3.html", "ex4.html"}
	for _, name := range fNames {
		f, err := os.Open(name)
		if err != nil {
			fmt.Printf("file open failed for file: %s\n", err)
			return
		}
		defer f.Close()
		output := linkparser.Parse(f)
		fmt.Printf("\noutput for file: %s\n", name)
		for idx, link := range output {
			fmt.Printf("%d. Href:%s \t Text: %s\n", idx, link.Href, link.Text)
		}
	}
}
