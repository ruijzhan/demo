package main

import "time"

func main() {
	runHello()
	runPubsubExample()
	runHttpGrpc()

	time.Sleep(time.Hour)
}
