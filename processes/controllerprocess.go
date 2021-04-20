package main

import (
	"cryptocurrency-project/ipaddresses"
	"net"
	"strings"
)

func main() {
	// Listen to specified TCP port of controller.
	controllerAddress := ipaddresses.GetController()
	port := strings.Split(controllerAddress, ":")[1]
	listener, err := net.Listen("tcp", port)

	// Do not proceed until all nodes and clients send their startup message.
	WaitForStartUps(listener)
	// Tell clients to start sending requests.
	StartClients()
	// Wait 30 seconds. Choose a random node to propose, if the node does not have transactions to propose,
	// choose another node. One minute after each node gets a proposal and tells the controller, send another proposal.
	ChooseNode()
}

