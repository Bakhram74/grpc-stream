package main

import (
	pb "grpc-demo/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloBydirectionalStreaming(stream pb.GreetService_SayHelloBydirectionalStreamingServer) error {
	log.Println("server stream started")
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		log.Printf("got req %s", req.Name)
		resp := &pb.HelloResponse{
			Message: "Hello" + req.Name,
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}
