package main

import (
	"fmt"
	"os"

	"github.com/codebasky/golang-examples/sitemap"
)

func main() {
	home := "https://www.calhoun.io"
	b := sitemap.New(home, 2)
	links, err := b.Build()
	if err != nil {
		fmt.Printf("site map building error : %s", err)
		return
	}
	sitemap.Encode(links, os.Stdout)
}
