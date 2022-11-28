package main

import (
	"time"

	"github.com/ruijzhan/demo/http/framework"
	"github.com/ruijzhan/demo/http/framework/middleware"
)

func registerRouter(core *framework.Core) {
	core.Get("/user/login", middleware.Test3(), UserLoginController)

	subjectApi := core.Group("/subject")
	{
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", middleware.Test3(), SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}

	core.Get("/foo", framework.TimeoutHandler(FooControllerHandler2, time.Second))
}
