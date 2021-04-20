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

	// Connect to specified tcp port of client.
	clientAddress := ipaddresses.GetClients()
	port := strings.Split(clientAddress[clientIndex], ":")[1]
	listener, err := net.Listen("tcp", port)

	// Store connection to the controller and the miners.



	// Send startup message to controller.
	client.SendStartup(clientAddress)

	// Wait for signal from controller.
	WaitForController(listener)

	// Periodically send out hard-coded clientTransactionRequest objects to nodes.
	SendRequest()
}
