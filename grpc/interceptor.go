package main

import (
	context "context"
	"log"

	"google.golang.org/grpc"
)

func unaryLogFilter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("unary filter:", info)
	return handler(ctx, req)
}

func unaryRecoverFilter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("unary recover:", err)
		}
	}()
	return handler(ctx, req)
}

func streamLogFilter(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("stream filter:", info)
	return handler(srv, ss)
}

func streamRecoverFilter(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println("stream recover:", err)
		}
	}()
	return handler(srv, ss)
}

var filters = []grpc.ServerOption{
	grpc.ChainUnaryInterceptor(unaryLogFilter, unaryRecoverFilter),
	grpc.ChainStreamInterceptor(streamLogFilter, streamRecoverFilter),
}
