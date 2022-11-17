package main

import (
	"io"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
}

func main() {

	defer func(f int) {
		log.SetFlags(f)
		log.Printf("%s", "hehe")
	}(log.Flags())

	log.SetPrefix("Login: ")
	log.Printf("%s", "haha")

	logger := log.New(io.MultiWriter(os.Stdout), "", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Writer()
	logger.Printf("%s", "haha")
}
