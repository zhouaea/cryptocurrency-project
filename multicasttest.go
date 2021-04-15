package main

import (
	"bufio"
	"cryptocurrency-project/application"
	"cryptocurrency-project/initialization"
	"os"
)

func main() {
	// Read information from command line and configuration file to obtain the node the process should act as
	// and a list of valid nodes to send messages to.
	node, nodes := initialization.InitializeNode()

	// Set up TCP listening for process, continuously listening for messages and printing them out.
	go node.ReceiveMessages()

	// Continuously scan user input for instructions on what message to send all nodes.
	// Send message to all nodes.
	scanner := bufio.NewScanner(os.Stdin)
	for {
		message := application.ReadCommands(scanner, node, nodes)
		go node.MulticastSendSimulatedDelay(message, nodes)
	}
}
