package main

import (
	"io"
	"log"

	pb "grpc-go/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {

	var messages []string

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}

		if err != nil {
			return err
		}

		log.Printf("Got Request with names: %v", req.name)
		messages = append(messages, "Hello"+req.Name)
	}
}
