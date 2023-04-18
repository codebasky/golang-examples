package pipeline

import "sync"

func Generator(done <-chan struct{}, num ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, item := range num {
			select {
			case out <- item:
			case <-done:
				return
			}
		}
	}()
	return out
}

func Processor(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for item := range in {
			select {
			case out <- item * item:
			case <-done:
				return
			}
		}
	}()
	return out
}

func Merger(done <-chan struct{}, chs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	extract := func(c <-chan int) {
		defer wg.Done()
		for item := range c {
			select {
			case out <- item:
			case <-done:
				return
			}
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
