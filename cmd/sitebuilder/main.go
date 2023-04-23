package main

import (
	"fmt"

	"github.com/codebasky/golang-examples/sitemap"
)

func main() {
	b := sitemap.New("https://www.calhoun.io", 4)
	output, err := b.Build()
	if err != nil {
		fmt.Printf("site map building error : %s", err)
		return
	}
	for link := range output {
		fmt.Println(link)
	}
}
