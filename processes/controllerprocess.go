package main

import (
	"cryptocurrency-project/controller"
	"cryptocurrency-project/errorchecker"
	"cryptocurrency-project/ipaddresses"
	"net"
	"strings"
	"time"
)

func main() {
	// Listen to specified TCP port of controller.
	controllerAddress := ipaddresses.GetController()
	port := strings.Split(controllerAddress, ":")[1]
	listener, err := net.Listen("tcp", port)
	errorchecker.CheckError(err)

	// Do not proceed until all nodes and clients send their startup message.
	nodeConnections := controller.WaitForStartUps(listener)

	// Tell clients to start sending requests.
	controller.StartClients(nodeConnections)

	// Wait 30 seconds. Choose a random node to propose, if the node does not have transactions to propose,
	// choose another node. One minute after each node gets a proposal and tells the controller, send another proposal.
	time.Sleep(10000)
	controller.ChooseNode()
}

