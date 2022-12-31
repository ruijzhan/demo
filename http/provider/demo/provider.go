package demo

import (
	"fmt"

	"github.com/ruijzhan/demo/http/framework"
)

type DemoServiceProvider struct {
}

func (sp *DemoServiceProvider) Name() string {
	return Key
}

func (sp *DemoServiceProvider) Reginster(_ framework.Container) framework.NewInstance {
	return NewDemoService
}

func (sp *DemoServiceProvider) Boot(_ framework.Container) error {
	fmt.Println("demo service boot")
	return nil
}

func (sp *DemoServiceProvider) IsDefer() bool {
	return true
}

func (sp *DemoServiceProvider) Params(c framework.Container) []interface{} {
	return []any{c}
}
