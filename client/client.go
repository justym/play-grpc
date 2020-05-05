package main

import (
	"context"
	"log"

	"github.com/justym/play-grpc/service"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("[Failed to connect] %+v", err)
	}
	defer conn.Close()

	client := service.NewMessageServiceClient(conn)

	msg := service.Message{
		Body: "Hello gRPC from Clinet !",
	}

	res, err := client.SayMessage(context.Background(), &msg)
	if err != nil {
		log.Fatalf("[Response Error] %+v", err)
	}

	log.Printf("[Received Message]: %s", res.Body)
}
