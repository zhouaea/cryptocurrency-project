package main

import (
	"cryptocurrency-project/user"
	"cryptocurrency-project/errorchecker"
	"cryptocurrency-project/ipaddresses"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read command line for user id to act as.
	if len(os.Args) != 2 {
		fmt.Println("Format: go run userprocess.go <user_index> ")
	}
	userIndex, err := strconv.Atoi(os.Args[1])
	errorchecker.CheckError(err)

	// Connect to specified tcp port of user.
	userAddress := ipaddresses.GetUsers()
	port := strings.Split(userAddress[userIndex], ":")[1]
	listener, err := net.Listen("tcp", port)

	// Store connection to the controller and the miners.



	// Send startup message to controller.
	user.SendStartup(userAddress)

	// Wait for signal from controller.
	WaitForController(listener)

	// Periodically send out hard-coded userTransactionRequest objects to nodes.
	SendRequest()
}
