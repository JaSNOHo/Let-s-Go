// Inspired by the Golang tutorial code for gRPC https://grpc.io/docs/languages/go/quickstart/
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

	pb "google.golang.org/grpc"
)

type Participant struct {
	name             string
	lamportTimestamp int
}

// unsure if we need default name, when we technically REQUIRE a name
const (
	defaultName = "Anonymous"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name of chatter")
)

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//client := &pb.NewClient(conn) //get the protobuff source from imports
	//above 7 lines of code form the Golang gRPC example

	clientRequest()
}

// inspired by https://www.youtube.com/watch?v=WB37L7PjI5k
func clientRequest() {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		//above two lines taken from Golang official example

		scanner := bufio.NewScanner(os.Stdin)
		input := scanner.Text()

		if input == "login" {
			print("Please enter a username:")
			*name = scanner.Text()
		} else if input == "logout" {
			break
		} else {
			message := scanner.Text()
			r, err := c.GetMessage(ctx, &pb.ClientRequest{Name: *name})
			//above line with err modified slightly from Golang official example

			if err != nil {
				log.Fatalf("Could not send message: %v", err)
				//unsure what %v prints above - maybe name, maybe message
			}
			//above message mostly taken from Golang example
			log.Printf(*name + ": " + message)
			//print the message and name of person. Log takes care of timestamp

			time.Sleep(1)
			//sleep for 1 to aid flow of messages

			//do something with lamport timestamp here
		}
	}
}
