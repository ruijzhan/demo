package middleware

import (
	"log"
	"time"

	"github.com/ruijzhan/demo/http/framework/gin"
)

func Cost() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func(start time.Time) {
			end := time.Now()
			cost := end.Sub(start)
			log.Printf("%s %s %v", c.Request.Method, c.Request.RequestURI, cost)
		}(time.Now())

		c.Next()
	}
}
