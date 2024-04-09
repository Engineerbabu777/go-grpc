package main

import (
	"context"
	"log"
	"time"

	pb "grpc-go/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {

	log.Printf("Client streaming started...")

	stream, er := client.SayHelloClientStreaming(context.Background())

	if er != nil {
		log.Fatalf("could not send names: %v", er)

	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("error sending request: %v", err)
		}
		log.Printf("Sent request with name: %v", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	log.Printf("client streaming finished")

	if err != nil {
		log.Fatalf("error while receiving %v", err)
	}

	log.Printf("%v", res.Messages)
}
