package miner

import (
	"cryptocurrency-project/ipaddresses"
	"cryptocurrency-project/message"
	"cryptocurrency-project/tcp"
	"fmt"
	"net"
	"os"
)

// SendStartup sends a start up message to the controller
func SendStartup(minerID int) net.Conn {
	message := message.startupMessage{"READY", "Miner " + string(minerID)}
	controllerChannel, err := net.Dial("tcp", ipaddresses.GetController())
	if err != nil {
		fmt.Println("Controller has not started up yet. Try again after booting up controller.")
		os.Exit(1)
	}
	tcp.Encode(controllerChannel, message)

	return controllerChannel
}
