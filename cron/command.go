package main

import "fmt"

func runEvery5s() {
	fmt.Println("Run every 5s")
}

type job func()

func (ff job) Run() {
	ff()
}
