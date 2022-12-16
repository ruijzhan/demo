package main

import (
	"time"

	"github.com/ruijzhan/demo/http/framework/gin"
	"github.com/ruijzhan/demo/http/framework/middleware"
)

func registerRouter(core *gin.Engine) {
	core.GET("/user/login", middleware.Timeout(time.Second), UserLoginController)

	subjectApi := core.Group("/subject")
	{
		subjectApi.DELETE("/:id", SubjectDelController)
		subjectApi.PUT("/:id", SubjectUpdateController)
		subjectApi.GET("/:id", middleware.Test3(), SubjectGetController)
		subjectApi.GET("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.GET("/name", SubjectNameController)
		}
	}

	// core.Get("/foo", framework.TimeoutHandler(FooControllerHandler2, time.Second))
}
