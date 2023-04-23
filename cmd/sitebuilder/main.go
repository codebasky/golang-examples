package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/codebasky/golang-examples/sitemap"
)

func main() {
	var domain string
	var depth int
	flag.StringVar(&domain, "domain", "https://www.calhoun.io", "website to map")
	flag.IntVar(&depth, "depth", 3, "depth of link parsing")
	flag.Parse()
	b := sitemap.New(domain, depth)
	links, err := b.Build()
	if err != nil {
		fmt.Printf("site map building error : %s", err)
		return
	}
	sitemap.Encode(links, os.Stdout)
}
