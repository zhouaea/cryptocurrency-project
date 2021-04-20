package main

import (
	"cryptocurrency-project/client"
	"cryptocurrency-project/errorchecker"
	"cryptocurrency-project/ipaddresses"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read command line for client id to act as.
	if len(os.Args) != 2 {
		fmt.Println("Format: go run clientprocess.go <client_index> ")
	}
	clientIndex, err := strconv.Atoi(os.Args[1])
	errorchecker.CheckError(err)

	// TODO I think clients actually do not need to listen to tcp ports because they will only have to communicate with the controller.
	// Connect to specified tcp port of client.
	clientAddresses := ipaddresses.GetClients()
	port := strings.Split(clientAddresses[clientIndex], ":")[1]
	listener, err := net.Listen("tcp", port)

	// Send startup message to controller and store the tcp channel from client to controller.
	controllerChannel := client.SendStartup(clientIndex)

	// Wait for signal from controller that signifies that all processes are running.
	client.WaitForController(controllerChannel)

	// Store connection to miners.

	// Periodically send out hard-coded clientTransactionRequest objects to nodes.
	client.SendRequest()
}
