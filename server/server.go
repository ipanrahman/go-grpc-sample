package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ipanrahman/go-grpc-sample/chat"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Go gRPC Sample!")
	listen, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}

	grpcServer := grpc.NewServer()

	var s chat.Server

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
