package main

import (
	"context"
	pb "grpc-go/proto"
)
func (s *helloServer) SayHello(ctx context.context, req *pb.NoParam) (*pb.HelloResponse, error) {

	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
