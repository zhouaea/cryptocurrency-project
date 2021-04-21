package controller

import (
	"cryptocurrency-project/errorchecker"
	"cryptocurrency-project/ipaddresses"
	"cryptocurrency-project/message"
	"cryptocurrency-project/tcp"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"
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
	startupMessage := new(message.Message)
	tcp.Decode(channel, startupMessage)

	// Associate the role of a process with the tcp channel that connects to the process.
	nodeConnections[startupMessage.Sender] = channel
	fmt.Printf("Received '%s' message from %s\n", startupMessage.Message, startupMessage.Sender)

	wg.Done()
}

// StartClients sends a PROCEED message via a string to every client via goroutines.
func StartClients(nodeConnections map[string]net.Conn) {
	numberOfClients := len(ipaddresses.GetClients())
	// NOTE: A full message object with a sender field is not required, since the clients will know that messages going
	// through this particular connection must come from the controller.
	startupMessage := "PROCEED"

	// Send the client a start up message in the form of a string.
	for i := 0; i < numberOfClients; i++ {
		go tcp.Encode(nodeConnections["Client " + string(i)], &startupMessage)
	}
}

// ChooseNode choose a random node to propose via a message in the form of a string, if the node does not have transactions to propose,
// choose another node. Once all miners have accepted or rejected a proposal
func ChooseNode(delayInterval int) {
	// Create and record connections to each miner.
	minerChannels := make(map[int]net.Conn)

	minerAddresses := ipaddresses.GetMiners()
	numberOfMiners := len(minerAddresses)
	for i := 0; i < numberOfMiners; i++ {
		minerChannel, err := net.Dial("tcp", minerAddresses[i])
		errorchecker.CheckError(err)
		minerChannels[i] = minerChannel
	}

	// Create a MINING SUCCESSFUL message (string) to send to all chosen miners.
	proposeMessage := "MINING SUCCESSFUL"
	
	for {
		// Randomly choose a miner.
		rand.Seed(time.Now().UnixNano())
		minerChosenIndex := rand.Intn(numberOfMiners)
		
		// Send a MINING SUCCESSFUL message to the randomly chosen miner.
		go tcp.Encode(minerChannels[minerChosenIndex], proposeMessage)
		time.Sleep(time.Duration(delayInterval) * time.Millisecond)
	}
}