package main

import (
	pb "grpc-demo/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var message []string
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&pb.MessageList{Messages: message})
			}
			return err
		}
		log.Printf("Got request with name %s", req.Name)
		message = append(message, "Hello", req.Name)
	}
}
