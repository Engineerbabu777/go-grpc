package main

import (
	"log"
	pb "grpc-go/proto"

	"google.golang.org/grpc"
)


const port = ":8080";

func main(){
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err!= nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn);

	names := &pb.NamesList{
		Names: []string{"Akhil","Alice","Bob"}
	}

	// callSayHello()
	// callSayHelloServerStream(client, names)
	callSayHelloClientStream(client, names)
}