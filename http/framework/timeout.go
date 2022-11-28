package framework

import (
	"context"
	"fmt"
	"log"
	"time"
)

func TimeoutHandler(f ControllerHandler, d time.Duration) ControllerHandler {
	return func(c *Context) error {

		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

		chFinish := make(chan struct{}, 1)
		chPanic := make(chan any, 1)

		go func() {
			defer func() {
				if p := recover(); p != nil {
					chPanic <- p
				}
			}()

			f(c)

			chFinish <- struct{}{}
		}()

		select {
		case p := <-chPanic:
			log.Println(p)
			c.responseWriter.WriteHeader(500)
		case <-durationCtx.Done():
			c.responseWriter.Write([]byte("time out"))
			c.SetHasTimeout()
		case <-chFinish:
			fmt.Println("finished")
		}

		return nil
	}
}

func Timeout(d time.Duration) ControllerHandler {
	return func(c *Context) error {
		return nil
	}
}
