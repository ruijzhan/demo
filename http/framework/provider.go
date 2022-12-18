package framework

type NewInstance func(...any) (any, error)

type ServiceProvider interface {
	Name() string

	Reginster(Container) NewInstance

	Boot(Container) error

	IsDefer() bool

	Params(Container) []interface{}
}
