package main

import (
	"fmt"

	"github.com/spf13/cast"
)

func main() {
	fmt.Println(cast.ToUint(8.1))
	_, err := cast.ToDurationE("1xx")
	if err != nil {
		panic(err)
	}
}
