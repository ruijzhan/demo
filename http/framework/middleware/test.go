package middleware

import (
	"fmt"

	"github.com/ruijzhan/demo/http/framework/gin"
)

func Test1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware pre Test1")
		c.Next()
		fmt.Println("middleware post Test1")
	}
}

func Test2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware pre Test2")
		c.Next()
		fmt.Println("middleware post Test2")
	}
}
func Test3() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware pre Test3")
		c.Next()
		fmt.Println("middleware post Test3")
	}
}
