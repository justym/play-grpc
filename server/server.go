package main

import (
	"log"
	"net"

	"github.com/justym/play-grpc/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("[Failed to listen on port]: %+v", err)
	}

	grpcService := service.Server{}
	grpcServer := grpc.NewServer()
	service.RegisterMessageServiceServer(grpcServer, &grpcService)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("[Failed to serve gRPC server]: %+v", err)
	}
}
