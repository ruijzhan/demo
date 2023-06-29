package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	myprc "github.com/ruijzhan/demo/rpc"
)

// RegisterHelloService registers the implementation of the HelloServiceInterface
// as the HelloServiceName with the rpc server
//
// service: the implementation of the HelloServiceInterface to be registered.
// returns an error if there are any issues during registration.
func RegisterHelloService(service myprc.HelloServiceInterface) error {
	return rpc.RegisterName(myprc.HelloServiceName, service)
}

type HelloService struct{}

// Hello is a function that takes a string parameter named request and a pointer to a string named reply.
// It returns an error.
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello: " + request
	return nil
}

func main() {
	RegisterHelloService(&HelloService{})

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		rpc.ServeRequest(jsonrpc.NewServerCodec(struct {
			io.Writer
			io.ReadCloser
		}{w, r.Body}))
	})

	http.ListenAndServe(":1234", nil)
}
