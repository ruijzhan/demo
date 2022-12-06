package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ruijzhan/demo/http/framework"
)

func FooControllerHandler(c *framework.Context) error {

	dc, cancel := context.WithTimeout(c.BaseContext(), time.Duration(time.Second))

	defer cancel()

	finish := make(chan struct{}, 1)
	panicChan := make(chan any, 1)

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		time.Sleep(10 * time.Second)
		c.SetStatus(200).Json("ok")
		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.SetStatus(500).Json(p)
	case <-dc.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()

		c.SetStatus(500).Json("timeout")

		c.SetHasTimeout()
	case <-finish:
		fmt.Println("finished")
	}

	return nil
}

func FooControllerHandler2(c *framework.Context) error {
	time.Sleep(10 * time.Second)
	// c.Json(200, "ok")
	return nil
}
