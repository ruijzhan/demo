package main

import (
	"context"
	fmt "fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"github.com/moby/pubsub"
	pb "github.com/ruijzhan/demo/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type PubsubService struct {
	pub *pubsub.Publisher
}

// NewPubsubService creates a new instance of PubsubService.
//
// No parameters.
// Returns a pointer to PubsubService.
func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

var _ pb.PubsubServiceServer = &PubsubService{}

// Publish is a method of the PubsubService struct that publishes a message.
//
// It takes a context.Context object and a *String object as input parameters.
// It returns a *String object and an error object as output parameters.
func (p *PubsubService) Publish(ctx context.Context, in *pb.String) (*pb.String, error) {
	p.pub.Publish(in.GetValue())
	return &pb.String{}, nil
}

// Subscribe handles subscribing to a topic and streaming messages.
//
// The function takes a String argument and a PubsubService_SubscribeServer
// stream parameter. It returns an error.
func (p *PubsubService) Subscribe(arg *pb.String, stream pb.PubsubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&pb.String{Value: v.(string)}); err != nil {
			return err
		}
	}

	return nil
}

func runPubsubExample() {
	stopCh := make(chan struct{})
	defer close(stopCh)
	runPubsubServer(stopCh)

	time.Sleep(time.Second)

	go runSubClient(stopCh)

	time.Sleep(time.Second)
	runPubClient()

	time.Sleep(time.Second)
}

func runPubsubServer(stopCh <-chan struct{}) {
	server := grpc.NewServer(filters...)
	pb.RegisterPubsubServiceServer(server, NewPubsubService())

	reflection.Register(server)

	lis, err := net.Listen("tcp", ":1235")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go server.Serve(lis)

	go func() {
		<-stopCh
		lis.Close()
	}()

}

func runSubClient(stopCh <-chan struct{}) {
	conn, err := grpc.Dial(":1235", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewPubsubServiceClient(conn)
	ctx := context.Background()

	topics := []string{
		"golang:",
		"docker:",
	}

	for _, topic := range topics {
		if err := subscribe(ctx, client, topic); err != nil {
			log.Fatal(err)
		}
	}

	<-stopCh
}

func subscribe(ctx context.Context, client pb.PubsubServiceClient, topic string) error {
	stream, err := client.Subscribe(ctx, &pb.String{Value: topic})
	if err != nil {
		return err
	}

	go func() {
		for {
			reply, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					log.Println(err)
				}
				break
			}
			fmt.Println(reply.GetValue())
		}
	}()

	return nil
}

func runPubClient() {
	conn, err := grpc.Dial(":1235", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewPubsubServiceClient(conn)

	msgs := []string{
		"golang: Hello Go",
		"docker: Hello Docker",
	}

	for _, msg := range msgs {
		if err := publish(client, msg); err != nil {
			log.Fatal(err)
		}
	}

}

func publish(client pb.PubsubServiceClient, msg string) error {
	_, err := client.Publish(context.Background(), &pb.String{Value: msg})
	return err
}
