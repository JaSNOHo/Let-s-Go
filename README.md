# Let-s-Go

a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?
We utilize channels to simulate the 3-way-handshake artificially. Our channel represents the communication between client and server, where they take turns sending+receiving information from each other. All data is contained and maintained by the communication channel, but manipulated by client and server (whom are represented by each of their functions, and would be different machines in the real world). We are using a struct as a container for the numbers that are sent between the client and the server. 

b) Does your implementation use threads or processes? Why is it not realistic to use threads?
It is not realistic to use threads because they communicate locally on the target computer. A real 3-way handshake is a communication between different machines. Using threads is just a simulation.
We use threads in the form of go routines.

c) In case the network changes the order in which messages are delivered, how would you handle message re-ordering?
We would have to change the sleep timers to handle the new ordering of messages. We are aware this is not a functional solution.

d) In case messages can be delayed or lost, how does your implementation handle message loss?
Our implementation does not handle loss of messages. In case of a delayed message, if it is delayed for longer than the sleep timer, our code continues and thus messes up the 3-way handshake.
We could have implementet if-statements that require messages to be received before code continues, or that breaks the process and connection upon, and returns an error, upon message loss.

e) Why is the 3-way handshake important?
The three-way handshake is used to establish a reliable connection between a client and a server. It ensures that both sides of the connection has synced up properly and can send data to eachother.