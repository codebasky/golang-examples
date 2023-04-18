package main

import (
	"fmt"

	"github.com/codebasky/golang-examples/pipeline"
)

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	done := make(chan struct{})
	defer close(done)
	in := pipeline.Generator(done, input...)
	p1 := pipeline.Processor(done, in)
	p2 := pipeline.Processor(done, in)
	result := pipeline.Merger(done, p1, p2)
	res := <-result
	fmt.Printf("square is %d\n", res)
}
