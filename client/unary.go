package main

import (
	"context"
	"log"
	"time"
)

func callSayHello(client pb.GreetserviceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := client.Sayhello(ctx, &pb.noParam())

	if err != nil {
		log.Fatalf("could not greet!")
	}
}
