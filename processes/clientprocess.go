package main

import (
	"cryptocurrency-project/client"
	"cryptocurrency-project/errorchecker"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read command line for client id to act as.
	if len(os.Args) != 2 {
		fmt.Println("Format: go run clientprocess.go <client_index> ")
	}
	clientIndex, err := strconv.Atoi(os.Args[1])
	errorchecker.CheckError(err)

	// NOTE: Clients actually do not need to listen to tcp ports because they will only have to communicate with the controller.

	// Send startup message to controller via READY message and store the communication channel.
	controllerChannel := client.SendStartup(clientIndex)

	// Wait for signal from controller that signifies that all processes are running.
	client.WaitForController(controllerChannel)

	// Periodically send out hard-coded clientTransactionRequest objects to nodes.
	client.SendRequest(10000)
}
