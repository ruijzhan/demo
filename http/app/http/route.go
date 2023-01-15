package http

import "github.com/ruijzhan/demo/http/framework/gin"

func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist/")

}
