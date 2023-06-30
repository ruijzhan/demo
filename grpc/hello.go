package main

import (
	context "context"
	fmt "fmt"
	"io"
	"log"
	"net"

	pb "github.com/ruijzhan/demo/grpc/proto"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type HelloServiceImpl struct{}

var _ pb.HelloServiceServer = &HelloServiceImpl{}

func (p *HelloServiceImpl) Hello(ctx context.Context, in *pb.String) (*pb.String, error) {
	reply := &pb.String{Value: "Hello: " + in.GetValue()}
	return reply, nil
}

func (p *HelloServiceImpl) Channel(stream pb.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		// 接收和发送的操作并不需要一一对应

		reply := &pb.String{Value: "Hello " + args.GetValue()}
		if err := stream.Send(reply); err != nil {
			return err
		}
	}
}

func runHelloServer(chStop <-chan struct{}) error {

	creds, err := credentials.NewServerTLSFromFile("certs/server.crt", "certs/server.key")
	if err != nil {
		return err
	}

	server := grpc.NewServer(append(filters, grpc.Creds(creds))...)
	pb.RegisterHelloServiceServer(server, &HelloServiceImpl{})

	reflection.Register(server)

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

func runHelloClient(addr string) error {

	creds, err := credentials.NewClientTLSFromFile("certs/server.crt", "server.grpc.io")
	if err != nil {
		return err
	}

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	stream, err := client.Channel(context.Background())
	if err != nil {
		return err
	}

	go func() {
		for i := 0; i < 3; i++ {
			msg := &pb.String{Value: fmt.Sprintf("World: %d", i)}
			if err := stream.Send(msg); err != nil {
				log.Fatal(err)
			}
		}
		// time.Sleep(time.Second)
		stream.CloseSend()
	}()

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		fmt.Println(reply.GetValue())
	}

	return nil
}

func runHello() {
	chStop := make(chan struct{})

	if err := runHelloServer(chStop); err != nil {
		log.Fatalln(err)
	}

	if err := runHelloClient("localhost:1234"); err != nil {
		log.Fatalln(err)
	}

	close(chStop)
}
