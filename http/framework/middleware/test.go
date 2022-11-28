package middleware

import (
	"fmt"

	"github.com/ruijzhan/demo/http/framework"
)

func Test1() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre Test1")
		c.Next()
		fmt.Println("middleware post Test1")
		return nil
	}
}

func Test2() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre Test2")
		c.Next()
		fmt.Println("middleware post Test2")
		return nil
	}
}
func Test3() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre Test3")
		c.Next()
		fmt.Println("middleware post Test3")
		return nil
	}
}
