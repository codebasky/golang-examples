package pipeline

import "sync"

func Generator(num ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, item := range num {
			out <- item
		}
	}()
	return out
}

func Processor(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for item := range in {
			out <- item * item
		}
	}()
	return out
}

func Merger(chs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	extract := func(c <-chan int) {
		defer wg.Done()
		for item := range c {
			out <- item
		}

	}
	wg.Add(len(chs))
	for _, ch := range chs {
		go extract(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
