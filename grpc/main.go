package main

import (
	context "context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, in *String) (*String, error) {
	reply := &String{Value: "Hello: " + in.GetValue()}
	return reply, nil
}

func main() {

	chStop := make(chan struct{})
	if err := runServer(chStop); err != nil {
		log.Fatalln(err)
	}

	_ = runClient()

	chStop <- struct{}{}
	time.Sleep(time.Second)
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
		fmt.Println("Server stopped.")
	}()
	return nil
}

func runClient() error {
	conn, err := grpc.Dial(":1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &String{Value: "world!"})
	if err != nil {
		return err
	}
	fmt.Println(reply.GetValue())

	return nil
}
