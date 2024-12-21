package main

import (
	"context"
	pb "grpc-demo/proto"
	"log"
	"time"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Println("server stream started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names")
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending %v", err)
		}
		log.Printf("sent req with name %s\n", name)
		time.Sleep(time.Second)
	}
	res, _ := stream.CloseAndRecv()
	
	log.Printf("server stream finished with values %v\n", res)
}
