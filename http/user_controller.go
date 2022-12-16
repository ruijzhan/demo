package main

import "github.com/ruijzhan/demo/http/framework/gin"

func UserLoginController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, UserLoginController")
}
