package main

import (
	context "context"
	"io"
	"net"

	grpc "google.golang.org/grpc"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, in *String) (*String, error) {
	reply := &String{Value: "Hello: " + in.GetValue()}
	return reply, nil
}

func (p *HelloServiceImpl) Channel(stream HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		// 接收和发送的操作并不需要一一对应

		reply := &String{Value: "Hello " + args.GetValue()}
		if err := stream.Send(reply); err != nil {
			return err
		}
	}
}

func runServer(chStop <-chan struct{}) error {
	server := grpc.NewServer()
	RegisterHelloServiceServer(server, &HelloServiceImpl{})

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		return err
	}

	go server.Serve(listener)
	go func() {
		<-chStop
		listener.Close()
		// fmt.Println("Server stopped.")
	}()
	return nil
}
