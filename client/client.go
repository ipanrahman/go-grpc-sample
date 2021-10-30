package main

import (
	"context"
	"log"
	"time"

	"github.com/ipanrahman/go-grpc-sample/chat"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did't connect: %v", err)
	}
	defer closeConn(conn)

	c := chat.NewChatServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &chat.Message{
		Body: "Hello, gRPC sample",
	})
	if err != nil {
		log.Panicf("error: %v", err)
	}
	log.Printf("info: %s", r.GetBody())
}

func closeConn(conn *grpc.ClientConn) {
	err := conn.Close()
	if err != nil {
		log.Fatalf("error:%v", err)
	}
}
