package main

import (
	"fmt"

	"github.com/codebasky/golang-examples/pipeline"
)

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	in := pipeline.Generator(input...)
	p1 := pipeline.Processor(in)
	p2 := pipeline.Processor(in)
	result := pipeline.Merger(p1, p2)
	for res := range result {
		fmt.Printf("square is %d\n", res)
	}
}
