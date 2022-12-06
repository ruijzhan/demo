package middleware

import (
	"context"
	"time"

	"github.com/ruijzhan/demo/http/framework"
)

func Timeout(d time.Duration) framework.ControllerHandler {
	return func(c *framework.Context) error {

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
		case <-dc.Done():
			c.SetStatus(500).Json("Time out")
			c.SetHasTimeout()
		case p := <-chPanic:
			panic(p)
		}

		return nil
	}
}
