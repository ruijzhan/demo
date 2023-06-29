package main

import (
	"fmt"
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

// NewHelloServiceClient dials a HelloServiceClient using the given network and address.
//
// network: The network protocol to use.
// address: The address of the service to dial.
// (*HelloServiceClient, error): Returns a client to the HelloService and an error.
func NewHelloServiceClient(network, address string) (*HelloServiceClient, error) {
	if network != "tcp" && network != "udp" {
		return nil, fmt.Errorf("network must be tcp or udp")
	}

	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cli := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return &HelloServiceClient{Client: cli}, nil
}

func main() {
	cli, err := NewHelloServiceClient("tcp", "localhost:1234")
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
