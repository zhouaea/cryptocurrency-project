package main

import (
	"MP1/initialization"
	"MP1/tcp"
)

func main() {
	// Read information from command line and configuration file to obtain the node the process should act as
	// and a list of valid nodes to send messages to.
	node, nodes := initialization.InitializeNode()

	// Set up TCP listening for process, continuously listening for messages and printing them out.
	go node.UnicastReceive()

	// Continuously scan user input for instructions on which node to send a message to.
	tcp.ExecuteCommands(node, nodes)
}
