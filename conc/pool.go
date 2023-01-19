package main

import (
	"fmt"

	"github.com/sourcegraph/conc/pool"
)

const (
	size = 1000
)

func handle(i int) {
	fmt.Println(i)
}

func tryPool() {
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = i
	}
	p := pool.New().WithMaxGoroutines(2)
	for _, i := range nums {
		n := i
		p.Go(func() {
			handle(n)
		})

	}
	p.Wait()
}

func tryReadChanel() {
	var s = make(chan int)

	go func() {
		defer close(s)
		for i := 0; i < size; i++ {
			s <- i
		}
	}()

	processStream(s)
}

func processStream(stream chan int) {
	p := pool.New().WithMaxGoroutines(10)
	for elem := range stream {
		elem := elem
		p.Go(func() {
			handle(elem)
		})
	}
	p.Wait()
}
