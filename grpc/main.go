package main

import (
	context "context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func runClient() error {
	conn, err := grpc.Dial(":1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := NewHelloServiceClient(conn)
	stream, err := client.Channel(context.Background())
	if err != nil {
		return err
	}

	go func() {
		for i := 0; i < 100000; i++ {
			msg := &String{Value: fmt.Sprintf("World: %d", i)}
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

func main() {

	chStop := make(chan struct{})
	if err := runServer(chStop); err != nil {
		log.Fatalln(err)
	}

	_ = runClient()

	close(chStop)

	time.Sleep(time.Second)
}
