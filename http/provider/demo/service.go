package demo

import (
	"fmt"

	"github.com/ruijzhan/demo/http/framework"
)

type DemoService struct {
	Service

	c framework.Container
}

func (s *DemoService) GetFoo() Foo {
	return Foo{
		Name: "I am Foo",
	}

}

func NewDemoService(params ...any) (any, error) {
	c := params[0].(framework.Container)

	fmt.Println("new demo service")
	return &DemoService{c: c}, nil
}
