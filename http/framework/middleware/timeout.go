package middleware

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ruijzhan/demo/http/framework/gin"
)

func Timeout(d time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {

		dc, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

		chFinish := make(chan struct{}, 1)
		chPanic := make(chan any, 1)

		go func() {
			defer func() {
				if p := recover(); p != nil {
					chPanic <- p
				}
			}()
			c.Next()
			chFinish <- struct{}{}
		}()

		select {
		case <-chFinish:
			fmt.Println("finish")
		case <-dc.Done():
			c.ISetStatus(500).IJson("Time out")
			// c.SetHasTimeout()
		case p := <-chPanic:
			c.ISetStatus(500).IJson("Time out")
			log.Println(p)
		}

	}
}
