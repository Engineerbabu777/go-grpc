package main

import (
	"log"
	"net"

	pb "grpc-go/proto"
	"google.golang.org/grpc"
)

type helloServer struct {
	pb.GreetServiceServer
}
var port = ":8080"

func main() {

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
