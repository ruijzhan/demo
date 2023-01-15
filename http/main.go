package main

import (
	"github.com/ruijzhan/demo/http/framework"
	"github.com/ruijzhan/demo/http/framework/provider/app"
)

// _ "net/http/pprof"

func main() {

	container := framework.NewMyContainer()

	container.Bind(&app.MyAppProvider{})

}
