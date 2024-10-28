// Inspired by the Golang tutorial code for gRPC https://grpc.io/docs/languages/go/quickstart/
// Inspired by https://www.youtube.com/watch?v=WB37L7PjI5k

package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "Let-s-Go/program"
)

type Participant struct {
	name             string
	lamportTimestamp int64
	stream           *pb.ChittyChat_ConnectClientClient
}

const (
	defaultName = "Anonymous"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)
var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name of chatter")
)

func main() {
	//below code handles previously defined flags
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChittyChatClient(conn)

	stream, err := c.ConnectClient(context.Background(), &pb.ClientName{
		User: *name})
	if err != nil {
		log.Fatalf("Connection failed %v", err)
	}

	participant := &Participant{
		name:             defaultName,
		lamportTimestamp: 0,
		stream:           &stream,
	}

	participant.clientRequest(c)
}

func (p *Participant) clientRequest(ServerConn pb.ChittyChatClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for scanner.Scan() {
		input := scanner.Text()

		if input == "login" {
			print("Please enter a username:")
			*name = scanner.Text()
			/*r, err := ServerConn.SendMessageToProgram(ctx, &pb.SendMessage{User: *name, Message: *name + " has joined ChittyChat", Timestamp: 0})
				if err != nil {
					log.Fatalf("You're invisible", err)
				}
			log.Printf("Message: %s", r)

			p.lamportTimestamp += 1*/

		} else if input == "logout" {
			/*r, err := ServerConn.SendMessageToProgram(ctx, &pb.SendMessage{User: *name, Message: *name + " has left ChittyChat", Timestamp: 0})
				if err != nil {
					log.Fatalf("You're invisible", err)
				}
			log.Printf("Message: %s", r)

			p.lamportTimestamp += 1*/
			break
		} else {
			if len(input) > 128 {
				log.Printf("Keep your message to a maximum of 128 characters")
			} else {
				r, err := ServerConn.SendMessageToProgram(ctx, &pb.SendMessage{User: *name, Message: input, Timestamp: 0})
				if err != nil {
					log.Fatalf("could not send message: %v", err)
				}
				log.Printf(*name, " is sending a message...")

				log.Printf("Message: %s", r)

				p.lamportTimestamp += 1
				//incrementing participants lamporttimestamp

				/*ServerConn.SendMessageToProgram(context.Background(), &pb.SendMessage{
					User:      *name,
					Message:   input,
					Timestamp: p.lamportTimestamp,
				})*/
			}
		}
	}
}
