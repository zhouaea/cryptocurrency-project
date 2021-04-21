package miner

import (
	"cryptocurrency-project/ipaddresses"
	"cryptocurrency-project/message"
	"cryptocurrency-project/tcp"
	"fmt"
	"net"
	"os"
	"time"
)

// SendStartup sends a start up message to the controller
func SendStartup(minerID int) net.Conn {
	message := message.startupMessage{"READY", "Miner " + string(minerID)}
	controllerChannel, err := net.Dial("tcp", ipaddresses.GetController())
	if err != nil {
		fmt.Println("Controller has not started up yet. Try again after booting up controller.")
		os.Exit(1)
	}
	tcp.Encode(controllerChannel, message)

	return controllerChannel
}

// MiningProtocol performs 3 actions:
// 1. Upon reception of ClientTransactionRequest objects, add them to the transaction array.
// 2. Upon reception of MINING SUCCESSFUL message from controller, propose block with transactions in transaction array.
// 3. Upon receival of PROPOSAL message, verify that the transactions are valid with correct signatures. If so,
//    telog.AddBlock(). Tell the controller that a PROPOSAL message was verified or denied with PROPOSAL RECEIVED.
func MiningProtocol(listener net.Listener) {
	for {
		// Wait for a connection from a client to our TCP port and then set up a TCP channel with them.
		conn, err := listener.Accept()
		errorchecker.CheckError(err)
		fmt.Println("Connection to sender was successful!")

		// Handle client as a goroutine to be able to handle multiple clients at once.
		go handleClient(conn)
	}
}

// Upon reception of ClientTransactionRequest objects, add them to the transaction array.
if (message == "REQUEST") {
ReceiveClientTransactions()
}

// Upon reception of CHOSEN message, propose block with transactions in transaction array.
if (message == "CHOSEN") {
Propose()
}

// Upon receival of propose message, verify that the transactions are valid with correct signatures. If so, telog.AddBlock().
// Tell the controller that a propose message was received with PROPOSE RECEIVED.
if (message == "PROPOSE") {
VerifyPropose()
}

// handleClient reads a message sent by another client, printing the message as well as the sender id and time received.
func handleClient(conn net.Conn) {
	defer conn.Close()
	// Read and print message sent by other client through TCP channel.
	message := new(message.Message)
	Decode(conn, message)
	time := time.Now()
	fmt.Println("Message received!")
	fmt.Println("---------------")
	fmt.Printf("Received '%s' from process %d\nmessage sent at %s\nmessage received at %s\n", message.Message, message.SenderId, message.TimeSent.Format("01-02-2006 15:04:05"), time.Format("01-02-2006 15:04:05"))
	fmt.Println("---------------")
}

