package middleware

import (
	"log"
	"time"

	"github.com/ruijzhan/demo/http/framework"
)

func Cost() framework.ControllerHandler {
	return func(c *framework.Context) error {
		defer func(start time.Time) {
			end := time.Now()
			cost := end.Sub(start)
			log.Printf("%s %s %v", c.GetRequest().Method, c.GetRequest().RequestURI, cost)
		}(time.Now())

		c.Next()
		return nil
	}
}
