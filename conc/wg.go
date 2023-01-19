package main

import (
	"fmt"
	"time"

	"github.com/sourcegraph/conc"
	"github.com/spf13/cast"
)

// tryWg declared a conc.WG in begining. Then creates 2 goroutines uses its Go methoed
// wg.Wait() will be blocked till the spawned goroutines exit
func tryWg() {
	//conc.WG has two fields
	var wg conc.WaitGroup

	defer func() {
		if val := recover(); val != nil {
			fmt.Println("recovered from panic")
		}
	}()

	defer wg.Wait()

	spawn(&wg)

}

func spawn(wg *conc.WaitGroup) {
	wg.Go(func() {
		fmt.Println("g1")
	})

	wg.Go(func() {
		fmt.Println("g2")
		time.Sleep(cast.ToDuration("1s"))
	})

	wg.Go(func() {
		fmt.Println("g3")
		panic("g3 panics")
	})
}
