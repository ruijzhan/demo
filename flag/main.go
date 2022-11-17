package main

import (
	"flag"
	"fmt"
)

func main() {
	args := []string{"-intflag", "12", "-stringflag", "test"}

	var intflag int
	var boolflag bool
	var stringflag string

	fs := flag.NewFlagSet("my", flag.ContinueOnError)
	fs.IntVar(&intflag, "intflag", 0, "int flag value")
	fs.BoolVar(&boolflag, "boolflag", false, "bool flag value")
	fs.StringVar(&stringflag, "stringflag", "default", "string flag value")

	fs.Parse(args)

	fs.Visit(func(f *flag.Flag) {
		fmt.Println(f.Name, f.Value)
	})

	fmt.Println("=========================")
	fmt.Println("int flag:", intflag)
	fmt.Println("bool flag:", boolflag)
	fmt.Println("string flag:", stringflag)
}
