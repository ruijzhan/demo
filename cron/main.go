package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {

	mc := myCron{
		cron: cron.New(
			cron.WithParser(cron.NewParser(cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow|cron.Descriptor)),
			cron.WithLogger(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))),
		),
		cronSpecs: []spec{},
	}

	mc.addJob("*/5 * * * * *", runEvery5s)
	mc.addFunc("@every 5s", runEvery5s)

	mc.start()

	for _, e := range mc.cronSpecs {
		fmt.Println(e.id)
	}

	for {
		time.Sleep(time.Second)
	}

}
