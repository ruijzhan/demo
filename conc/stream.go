package main

import (
	"fmt"

	"github.com/sourcegraph/conc/stream"
)

func mapStream(in, out chan int, f func(int) int, concurent int) {

	p := stream.New().WithMaxGoroutines(concurent)
	for elem := range in {
		elem := elem
		p.Go(func() stream.Callback {
			res := f(elem)
			return func() {
				out <- res
			}
		})
	}
	p.Wait()
}

func tryStream() {
	var in = make(chan int)
	var out = make(chan int)
	var out2 = make(chan int)
	defer close(out2)

	f := func(i int) int {
		return i + 1
	}

	// sender must close chanel when finishing send data
	go func() {
		defer close(in)
		for i := 0; i < 10; i++ {
			in <- i
		}
	}()

	go func() {
		for i := range out2 {
			fmt.Println(i)
		}
	}()

	go func() {
		defer close(out)
		mapStream(in, out, f, 2)
	}()

	mapStream(out, out2, f, 3)
}
