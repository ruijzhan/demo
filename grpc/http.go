package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	pb "github.com/ruijzhan/demo/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func serveHttpGrpc() {

	creds, err := credentials.NewServerTLSFromFile("certs/server.crt", "certs/server.key")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterHelloServiceServer(grpcServer, &HelloServiceImpl{})

	reflection.Register(grpcServer)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})

	err = http.ListenAndServeTLS(":8080", "certs/server.crt", "certs/server.key", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor != 2 {
			mux.ServeHTTP(w, r)
			return
		}
		if strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
			return
		}

		mux.ServeHTTP(w, r)
	}))

	if err != nil {
		log.Fatalln(err)
	}
}

func httpClient() {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))
}

func grpcClient() {
	runHelloClient("localhost:8080")
}

func runHttpGrpc() {
	go serveHttpGrpc()
	time.Sleep(time.Second)
	httpClient()
	grpcClient()
}
