package main

import(
	"bufio"
	"cryptocurrency-project/application"
	"cryptocurrency-project/ipaddresses"
	"os"
)

func main() {
	// Read information from command line and configuration file to obtain the client the process should act as
	// and a list of valid nodes to send message to.
	node, nodes := ipaddresses.InitializeNode()

	// Set up TCP listening for process, continuously listening for message and printing them out.
	go node.ReceiveMessages()

	// Continuously scan client input for instructions on what transaction to send all nodes.
	// Send transaction to all nodes.
	scanner := bufio.NewScanner(os.Stdin)
	for {
		message := application.ReadCommands(scanner, node, nodes)
		go node.MulticastSendSimulatedDelay(message, nodes)
	}
}
