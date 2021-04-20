package controller

import (
	"cryptocurrency-project/errorchecker"
	"cryptocurrency-project/ipaddresses"
	"cryptocurrency-project/message"
	"cryptocurrency-project/tcp"
	"fmt"
	"net"
	"sync"
)

var requiredProcesses = 7

// WaitForStartUps will not exit until all nodes in the cryptocurrency simulation are connected to the controller.
func WaitForStartUps(listener net.Listener) map[string]net.Conn {
	nodeConnections := make(map[string]net.Conn)
	nodesConnected := 0
	var wg sync.WaitGroup

	for {
		// Wait for a connection from a client to our TCP port and then set up a TCP channel with them.
		channel, err := listener.Accept()
		errorchecker.CheckError(err)

		// Increment count of nodes connected.
		nodesConnected += 1
		fmt.Printf("Connection to process was successful! %i nodes connected.", nodesConnected)

		// Handle client as a goroutine to be able to handle multiple clients at once.
		wg.Add(1)
		go recordNodeConnections(wg, channel, nodeConnections)

		// Exit function once enough nodes are connected
		if nodesConnected >= requiredProcesses {
			// Wait for all connections to be recorded before proceeding.
			wg.Wait()
			return nodeConnections
		}
	}
}

// recordNodeConnections reads a startup message sent through a tcp channel and associates the role of a process with
// the tcp channel that connects to the process.
func recordNodeConnections(wg sync.WaitGroup, channel net.Conn, nodeConnections map[string]net.Conn) {
	// Read and print message sent by miner or client through a tcp channel.
	startupMessage := new(message.StartupMessage)
	tcp.Decode(channel, startupMessage)

	// Associate the role of a process with the tcp channel that connects to the process.
	nodeConnections[startupMessage.Sender] = channel
	fmt.Printf("Received '%s' message from %s\n", startupMessage.Message, startupMessage.Sender)

	wg.Done()
}

// StartClients sends a PROCEED message to every client via goroutines.
func StartClients(nodeConnections map[string]net.Conn) {
	numberOfClients := len(ipaddresses.GetClients())
	// A full message object with the sender is not required since the clients will know that messages going through
	// their particular connection must come from the controller.
	startupMessage := "PROCEED"

	for i := 0; i < numberOfClients; i++ {
		go sendClientStartup(nodeConnections["Client " + string(i)], &startupMessage)
	}
}

func sendClientStartup(clientConnection net.Conn, startupMessage *string) {
	tcp.Encode(clientConnection, *startupMessage)
}