package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	myrpc "github.com/ruijzhan/demo/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

var _ myrpc.HelloServiceInterface = (*HelloServiceClient)(nil)

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Call(myrpc.HelloServiceName+".Hello", request, reply)
}

// DialHelloService creates a new client to communicate with the HelloService using the given network and address.
//
// network: a string representing the network protocol to use (e.g. "tcp", "udp").
// address: a string representing the network address to connect to (e.g. "localhost:1234").
//
// *HelloServiceClient: a pointer to a new instance of HelloServiceClient.
// error: if there was an error dialing the service, it will be returned.
func DialHelloService(network, address string) (*HelloServiceClient, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	cli := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return &HelloServiceClient{Client: cli}, nil
}

func main() {
	cli, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("DialHelloService error:", err)
	}

	var reply string
	err = cli.Hello("world", &reply)
	if err != nil {
		log.Fatal("Hello error:", err)
	}
	defer cli.Close()

	log.Println(reply)
}
