package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"sort"

	"github.com/codebasky/golang-examples/pipeline"
)

func main() {
	fmt.Println("Digester Pipeline Example")
	if len(os.Args) != 2 {
		fmt.Println("Need path as input parameter")
		return
	}
	path := os.Args[1]
	results, ech := pipeline.MD5ALL(path)
	digest := make(map[string][md5.Size]byte)
	for result := range results {
		if result.Err != nil {
			fmt.Printf("Error %s on processing file %s", result.Err, result.Path)
			return
		}
		digest[result.Path] = result.Digest
	}
	fpath := []string{}
	for path := range digest {
		fpath = append(fpath, path)
	}
	sort.Strings(fpath)
	for _, path := range fpath {
		fmt.Printf("Digest for file: %s is %x\n", path, digest[path])
	}
	err := <-ech
	if err != nil {
		fmt.Printf("Error processing walk %s", err)
	}
}
