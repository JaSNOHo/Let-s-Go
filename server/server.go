package main

import (
	"context"
	"flag"
	"log"

	pb "google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

//above code from Golang gRPC example

type Server struct {
	pb.UnimplementedTimeAskServer
	name string
	port int
	//messages         []*ClientStream inplement struct for this
	LamportTimestamp int64
}

//above code from Golang gRPC example

func (s *server) GetMessage(_ context.Context, in *pb.ClientName) (*pb.SendMessage, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.SendMessage{Message: in.message + in.GetName()}, nil
}

//above code adapted from Golang gRPC example
