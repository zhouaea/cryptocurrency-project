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

	// NOTE: Clients actually do not need to listen to tcp ports because they will only have to communicate with the controller.
	// Connect to specified tcp port of miner.
	clientAddresses := ipaddresses.GetClients()
	port := strings.Split(clientAddresses[clientIndex], ":")[1]
	listener, err := net.Listen("tcp", port)

	// Send startup message to controller via READY message and store the communication channel.
	controllerChannel := client.SendStartup(clientIndex)

	// Wait for signal from controller that signifies that all processes are running.
	client.WaitForController(controllerChannel)

	// Wait for all miners to connect to client.
	minerChannels := client.ConnectToMiners(listener)

	// Periodically send out hard-coded clientTransactionRequest objects to nodes.
	client.SendRequest(minerChannels, 10000)
}
