# Let-s-Go

Our program runs on both the server and client.
In order to run this program (which is not fully implemented!
See report for details), you need to open two terminals.
- In the first terminal, navigate to the /server folder.
- Enter 'go run server.go' in the terminal, and observe that the logs tell
  you, that the server has connected and is listening.
- Allow this program to run in the background.

- In the second terminal, navigate to the /client folder.
- Enter 'go run client.go'.
- Enter 'login' to activate a prompt asking for a username.
- ---From here, the rest is not implemented, but should work as such:---
- After entering a username, message can be entered, and posted by way
  of the 'enter' button. Participant can keep sending messages, and will
  continuously receive messages from other Participants by way of the server.
- To leave the chatroom, a Participant only needs to enter 'logout'.

As the program is running on the localhost, in case client wants to chat with
several Participants, server still only needs to be started once, and then
the desired number of extra terminals corresponding to the number of
Participants, need to be opened and run the client main code.
