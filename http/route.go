package main

import "github.com/ruijzhan/demo/http/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
