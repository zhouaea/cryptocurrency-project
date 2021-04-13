package tcp

import (
	"cryptocurrency-project/errorchecker"
	"cryptocurrency-project/messages"
	"fmt"
	"net"
	"time"
)

// Configure a node to listen for tcp connections.
func (node Node) UnicastReceive() {
	// Listen to an unused TCP port on localhost.
	port := ":" + node.Port
	listener, err := net.Listen("tcp", port)
	errorchecker.CheckError(err)
		defer listener.Close()
	fmt.Println("Listening to tcp port " + port + " was successful!")
	fmt.Println("To send a message, type: send <destination_id_number> <message>")

	// Listen for TCP connections until the process is closed.
	for {
		// Wait for a connection from a client to our TCP port and then set up a TCP channel with them.
		conn, err := listener.Accept()
		errorchecker.CheckError(err)
		fmt.Println("Connection to client was successful!")

		// Handle client as a goroutine to be able to handle multiple clients at once.
		go handleClient(conn)
	}
	return
}

// handleClient reads a message sent by another node, printing the message as well as the sender id and time received.
func handleClient(conn net.Conn) {
	defer conn.Close()
	// Read and print message sent by client through TCP channel
	message := new(messages.Message)
	messages.Decode(conn, message)
	time := time.Now()
	fmt.Println("Message received!")
	fmt.Println("---------------")
	fmt.Printf("Received '%s' from process %d\nmessage sent at %s\nmessage received at %s\n", message.Message, message.SenderId, message.TimeSent.Format("01-02-2006 15:04:05"), time.Format("01-02-2006 15:04:05"))
	fmt.Println("---------------")
}
