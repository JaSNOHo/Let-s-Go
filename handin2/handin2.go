package main

import (
	"fmt"
	"sync"
	"time"
)

var comm = make(chan numbers, 1)

type numbers struct {
	clientSeq int
	serverSeq int
}

func main() {
	println("---ATTEMPTING CONNECTION BETWEEN CLIENT & SERVER---")

	seqAndAck := numbers{
		clientSeq: 0,
		serverSeq: 0,
	}

	comm <- seqAndAck
	var wait sync.WaitGroup
	wait.Add(2)
	go client(&seqAndAck, &wait)
	go server(&seqAndAck, &wait)
	wait.Wait()
	println("Connection is established with clientnumber: " + fmt.Sprint(seqAndAck.clientSeq) + " and servernumber: " + fmt.Sprint(seqAndAck.serverSeq))
}

func client(seqAndAck *numbers, wait *sync.WaitGroup) {
	//client to server
	<-comm
	seqAndAck.clientSeq = 100
	println("client sending clientnumber to server")
	comm <- *seqAndAck
	time.Sleep(time.Duration(2000) * time.Millisecond)

	//final acknowledgement
	<-comm
	seqAndAck.serverSeq++
	comm <- *seqAndAck
	println("client incrementing servernumber, sending final acknowledgement for connection")
	wait.Done()
}

func server(seqAndAck *numbers, wait *sync.WaitGroup) {
	time.Sleep(time.Duration(1000) * time.Millisecond)
	//server to client
	<-comm
	println("Server has recieved client number")
	seqAndAck.clientSeq++
	seqAndAck.serverSeq = 300
	println("Server incrementing clientnumber, sending it+servernumber back to client")
	comm <- *seqAndAck
	wait.Done()
}
