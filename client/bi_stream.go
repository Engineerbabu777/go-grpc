package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "grpc-go/proto"

)

func callHelloBidirectionalStream(client pb.GreeterClientService, names *pb.NamesList) {

	log.Printf("Calling HelloBidirectionalStream RPC...")

	stream, err := client.HelloBidirectionalStreaming(context.Background())

	if err != nil {
		log.Fatalf("could not send names %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while streaming %v: %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v: %v", err)
		}

		time.Sleep(2 * time.Second)
	}

	stream.CloseSend()
}
