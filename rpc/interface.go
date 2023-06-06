package rpc

const HelloServiceName = "HelloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}
