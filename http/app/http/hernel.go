package http

import "github.com/ruijzhan/demo/http/framework/gin"

func NewHttpEngine() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	return r, nil
}
