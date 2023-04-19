package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/codebasky/golang-examples/pipeline"
)

func main() {
	fmt.Println("Digester Pipeline Example")
	done := make(chan struct{})
	defer close(done)
	if len(os.Args) != 2 {
		fmt.Println("Need path as input parameter")
		return
	}
	path := os.Args[1]
	digest, err := pipeline.MD5ALL(done, path)
	fpath := []string{}
	for path := range digest {
		fpath = append(fpath, path)
	}
	sort.Strings(fpath)
	for _, path := range fpath {
		fmt.Printf("Digest for file: %s is %x\n", path, digest[path])
	}

	if err != nil {
		fmt.Printf("Error processing walk %s", err)
	}
}
