// Inspired by the Golang tutorial code for gRPC https://grpc.io/docs/languages/go/quickstart/
// inspired by https://www.youtube.com/watch?v=WB37L7PjI5k

package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "Let-s-Go/program"
)

type Participant struct {
	name             string
	lamportTimestamp int64
	stream           *pb.ChittyChat_ConnectClientClient
}

// unsure if we need default name, when we technically REQUIRE a name
const (
	defaultName = "Anonymous"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)
var (
	addr = flag.Int("addr", 50051, "the address to connect to")
	name = flag.String("name", defaultName, "Name of chatter")
)

func main() {
	//below line handles previously defined flags
	flag.Parse()

	ServerConn, _ := ConnectServer()
	stream, err := ServerConn.ConnectClient(context.Background(), &pb.ClientName{
		User: *name})
	if err != nil {
		log.Fatalf("Connection failed")
	}

	participant := &Participant{
		name:             defaultName,
		lamportTimestamp: 0,
		stream:           &stream,
	}

	go participant.clientRequest()

	//scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		if input == "login" {
			print("Please enter a username:")
			*name = scanner.Text()
		} else if input == "logout" {
			break
		} else {
			log.Printf(*name, " is sending a message...")

			participant.lamportTimestamp += 1
			//incrementing participants lamporttimestamp

			ServerConn.SendMessageToProgram(context.Background(), &pb.SendMessage{
				User:      *name,
				Message:   input,
				Timestamp: participant.lamportTimestamp,
			})
		}
	}
}

func (p *Participant) clientRequest() {
	for {
		message, err := (*p.stream).Recv()
		if err != nil {
			log.Fatalf("Failed to send message due to: %v", err)
		}

		if message.Timestamp > p.lamportTimestamp {
			p.lamportTimestamp = message.Timestamp + 1
		} else {
			p.lamportTimestamp += 1
		}
		//above 5 lines from someone elses code - can we use this?

		log.Printf(p.name, " has sent message: ", message.Message, " at time: ", p.lamportTimestamp)
	}
}

func ConnectServer() (pb.ChittyChatClient, error) {
	conn, err := grpc.NewClient("localhost:"+strconv.Itoa(*addr), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	} else {
		log.Printf("Connected succesfully to port: %v", *addr)
	}
	return pb.NewChittyChatClient(conn), nil
}
