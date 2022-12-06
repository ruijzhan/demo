package middleware

import (
	"github.com/ruijzhan/demo/http/framework"
)

func Recovery() framework.ControllerHandler {
	return func(c *framework.Context) error {
		defer func() {
			if err := recover(); err != nil {
				c.SetStatus(500).Json(err)
			}
		}()

		c.Next()
		return nil
	}
}
