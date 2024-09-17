package main

import (
	"strconv"
)

var comm = make(chan int, 2)

type numbers struct{
	clientSeq int;
	serverSep int;
}

	//SYN clientSeq
	//SYN-ACK clientSeq=clientAck OG serverSeq
	//ACK serverSeq(+1)=clientAck  OG clientAck=clientSeq
func main() {
	println("---ATTEMPTING CONNECTION BETWEEN CLIENT & SERVER---")
	
	seqAndAck := numbers{
		clientSeq : 100,
		serverSeq : 300,
	}

	clientSeq <- comm.seqAndAck.clientSeq + 1

	comm <- seqAndAck.clientSeq++
	println("Client tries to synchronize by sending sequence " + strconv.Itoa(clientSeq))
	
	clientSeq++
	comm <- clientSeq ++ && serverSeq
	println("Server has acknowledged sequence number from client")
	println("Servers tries to synchronize and sends acknowledge number " + strconv.Itoa(clientSeq) " and sequence number " + strconv.Itoa(serverSeq) " to client")
	
	serverSeq++
	finalAckNumber := <- comm
	println("Client sends acknowledgement back to server with acknowledge number " + strconv.Itoa(serverSeq) + " and sequence number " + strconv.Itoa(clientSeq))
	println("Final acknowledge data is received")
	println("---CONNECTION ESTABLISHED---")
	
	
//	println(finalAck)
}
