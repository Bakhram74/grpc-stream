package main

import (
	"context"
	"fmt"
	pb "grpc-demo/proto"
	"log"
	"time"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("client.SayHello %v", err)
	}
	fmt.Printf("client.SayHello %s", resp.Message)
}

