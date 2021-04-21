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

	nodeConnections := controller.WaitForStartUps(listener)

	controller.StartClients(nodeConnections)

	time.Sleep(time.Duration(10000) * time.Millisecond)
	controller.ChooseNode(60000)
}

