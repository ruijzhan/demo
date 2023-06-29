package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

func (p *Person) GetName() string {
	return p.Name
}

type MyInt int

func (i MyInt) ToString() string {
	return fmt.Sprintf("%d", i)
}

func main() {
	var v MyInt
	t := reflect.TypeOf(v)
	fmt.Println(t.NumMethod())
}
