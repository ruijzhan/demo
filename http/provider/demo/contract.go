package demo

const Key = "my:demo"

type Service interface {
	GetFoo() Foo
}

type Foo struct {
	Name string
}
