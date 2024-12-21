package main

import (
	"context"
	pb "grpc-demo/proto"
	"io"
	"log"
	"time"
)

func callSayHelloBiDirectioanlStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Println("client stream started")
	stream, err := client.SayHelloBydirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names %v", err)
	}

	waitCh := make(chan struct{})

	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalf("error while streaming %v", err)
			}
			log.Printf("got resp %s", resp.Message)
		}
		close(waitCh)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending %v", err)
		}
		time.Sleep(time.Second)
	}
	stream.CloseSend()
	<-waitCh
	log.Println("client stream finshed")
}
