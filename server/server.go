package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sync"

	pb "Let-s-Go/program"
	//"google.golang.org/grpc/encoding/proto"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

//above code from Golang gRPC example

type Participants struct {
	pName string
	stream *pb.ChittyChat_ConnectClientServer

}

type Server struct {
	pb.UnimplementedChittyChatServer
	port int
	particpants []*Participants
	mu   sync.Mutex
}

func main() {
	//below line handles previously defined flags
	flag.Parse()

	//initializer new server objects
	server := &Server{
		port * port,
		participants: []*Participants{},
		mu sync.Mutex,
	}

	//Sets server to listen for RPCs on its port
	go TurnOnServer()
}

//above code from Golang gRPC example

func (s *server) GetMessage(_ context.Context, in *pb.ClientName) (*pb.SendMessage, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.SendMessage{Message: in.message + in.GetName()}, nil
}

//above code adapted from Golang gRPC example

func TurnOnServer(server *Server) {

	//this method is taken from ChatGPt


	listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterChittyChatServer(grpcServer, server)

    log.Println("Chat server is running on :50051...")
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
	

	
	
}

// SendMessage()
// 1. catch message from client
// 2. decode message parameters
// 3.send message to public, broadcast?
func () SendMessage(string Username, string message) { //ctx context.Context, in pb.sendmessage)(*pb.messagePublished, error?)
	s.mu.Lock()
	s.LamportTimestamp++
	s.mu.Unlock()

	//log message that incluses users name and message content for tracking
	broadcastMsg := fmt.Sprintf("%s: %s", in.GetUser(), in.GetMessage())

	//call bc helper func, passing a new sendmsg object
	//object includes the sender's name, the msg content, and the current lmpts
	s.broadcast(&pb.SendMessage{
		User:      in.GetUser(),
		Message:   in.GetMessage(),
		Timestamp: s.LamportTimestamp,
	})

	//to server log for tracking
	log.Printf("Published message: %s", &broadcastMsg)

	//respond to sender w confirmation msg
	return &pb.MessagePublished{
		ServerName: "ChittyChatServer",
		Timestamp:  s.LamportTimestamp,
	}
}

// IsMessageSent()
func (s *Server) ReceiveBroadcast(empty *pb.Empty, stream pb.ChittyChat_RecieveBroadcastServer) error {
	clientCh := make(chan *pb.SendMessage, 10)

	//listen for msgs on cli chan and send to cli stream
	for msg := range clientCh {
		if err := stream.Send(msg); err != nil {
			return err
		}
	}
	return nil
}

// Disconnect()
func (s *Server) DisconnectUser(ctx context.Context, in *pb.ClientName) (*pb.LeaveResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	clientName := in.GetUser()
	s.LamportTimestamp++
	leaveMessage := fmt.Sprintf("Client %s left the chat at Lamport time %d", clientName, a.LamportTimestamp)

	//close ch and rm from map
	if ch, ok := s.clients[clientName]; ok {
		close(ch)
		delete(s.clients, clientName)
	}

	//need to broadcasst the leave message to remaining clients
	s.broadcast(&pb.SendMessage{User:"Server", Message: leaveMessage, Timestamp: s.LamportTimestamp})
	log.Printf(leaveMessage)
	return &pb.LeaveResponse{leaveMessage, Timestamo: s.LamportTimestamp}, nil
}

//MessagePublish()
//FindMessage()
//SendToStream()

/// User: mnfushnfoiew gji (timestamp)
