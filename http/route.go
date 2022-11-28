package main

import (
	"time"

	"github.com/ruijzhan/demo/http/framework"
)

func registerRouter(core *framework.Core) {
	core.Get("/user/login", framework.TimeoutHandler(UserLoginController, time.Second))

	subjectApi := core.Group("/subject")
	{
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}

	core.Get("/foo", framework.TimeoutHandler(FooControllerHandler2, time.Second))
}
