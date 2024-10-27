package main

import (
	"context"
	"flag"
	"log"
	"net"
	"sync"

	pb "Let-s-Go/program"

	"google.golang.org/grpc"
	//"google.golang.org/grpc/encoding/proto"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

//above code from Golang gRPC example

type Participants struct {
	pName  string
	stream *pb.ChittyChat_ConnectClientServer
}

type Server struct {
	pb.UnimplementedChittyChatServer
	port             int
	participants     []*Participants
	mu               sync.Mutex
	LamportTimestamp int64
}

func main() {
	//below line handles previously defined flags
	flag.Parse()

	//server :=grpc.TurnOnServer()

	server := &Server{
		port:             *port,
		participants:     []*Participants{},
		mu:               sync.Mutex{},
		LamportTimestamp: 0,
	}

	//Sets server to listen for RPCs on its port
	TurnOnServer(server)

}

func (s *Server) GetMessage(_ context.Context, in *pb.SendMessage) (*pb.SendMessage, error) {
	log.Printf("Received: %v", in.User)
	return &pb.SendMessage{Message: in.Message + in.User}, nil
}

//above code adapted from Golang gRPC example

func TurnOnServer(server *Server) {

	//this method is based on code provided by ChatGPT

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	} else {
		log.Printf("Now listening on: 50051")
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
/*func SendMessage(Username string, message string, s *Server, in *pb.SendMessage) *pb.MessageSent {
	//ctx context.Context, in pb.sendmessage)(*pb.messagePublished, error?)
	s.mu.Lock()
	s.LamportTimestamp++
	s.mu.Unlock()

	//log message that incluses users name and message content for tracking
	broadcastMsg := fmt.Sprintf("%s: %s", in.User, in.Message)

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
	return &pb.MessageSent{
		Server:    "ChittyChatServer",
		Timestamp: s.LamportTimestamp,
	}
}

// IsMessageSent()
// Previous: empty *pb.Empty and pb.ChittyChat_RecieveBroadcastServer as previous stream argument
func (s *Server) ReceiveBroadcast(stream pb.ChittyChat_ConnectClientServer) error {
	//previously: clientCh := make(chan *pb.SendMessage, 10)
	clientCh := make(chan *pb.MessagePublished, 10)

	//listen for msgs on cli chan and send to cli stream
	for msg := range clientCh {
		if err := stream.Send(msg); err != nil {
			return err
		}
	}
	return nil
}

// Disconnect()
func (s *Server) DisconnectUser(ctx context.Context, in *pb.ClientName) (*pb.SendMessage, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	clientName := in.GetUser()
	s.LamportTimestamp++
	leaveMessage := fmt.Sprintf("Client %s left the chat at Lamport time %d", clientName, s.LamportTimestamp)

	//close ch and rm from map

	//	if ch, ok := s.clients[clientName]; ok { //s.participants[pName]?

	if ch, ok := s.participants[pName]; ok {
		close(ch)
		delete(s.participants, s.participants[pName]) //s.clients, clientName
	}

	//need to broadcasst the leave message to remaining clients
	s.broadcast(&pb.SendMessage{User: "Server", Message: leaveMessage, Timestamp: s.LamportTimestamp})
	log.Printf(leaveMessage)
	//previously return &pb.LeaveResponse{leaveMessage, Timestamp: s.LamportTimestamp}, nil
	return &pb.SendMessage{User: clientName, Message: leaveMessage, Timestamp: s.LamportTimestamp}, nil
}

// MessagePublish()
func (s *Server) broadcast(msg *pb.SendMessage) {
	s.mu.Lock()
	defer s.mu.Unlock()

	//send to active clients
	for _, ch := range s.particpants {
		ch <- msg
	}
}
*/
//FindMessage()
//SendToStream()

/// User: mnfushnfoiew gji (timestamp)
