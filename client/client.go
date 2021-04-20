package client

import (
	"../esign"
	"crypto/ed25519"
	"cryptocurrency-project/ipaddresses"
	"cryptocurrency-project/message"
	"cryptocurrency-project/tcp"
	"fmt"
	"net"
	"os"
)

type client struct {
	Name       string
	Sk ed25519.PrivateKey
	Pk ed25519.PublicKey
}

type Clients struct {
	all []client
}

// Creates a new client
func (n *Clients) NewClient(name string) client {
	pk, sk := esign.GenerateKeyPair()

	new := client{
		Name: name,
		Sk: sk,
		Pk: pk,
	}

	n.all = append(n.all, new)
	return new
}

// SendStartup sends a start up message to the controller
func SendStartup(clientID int) net.Conn {
	message := message.StartupMessage{"READY", "Client " + string(clientID)}
	controllerChannel, err := net.Dial("tcp", ipaddresses.GetController())
	if err != nil {
		fmt.Println("Controller has not started up yet. Try again after booting up controller.")
		os.Exit(1)
	}
	tcp.Encode(controllerChannel, message)

	return controllerChannel
}

// WaitForController waits for a signal from the controller that all processes have started up.
func WaitForController(controllerChannel net.Conn) {
	for {
		// Wait for a message to be sent from the controller through the tcp channel
		var startupMessage string
		tcp.Decode(controllerChannel, startupMessage)

		// Exit function once controller sends a PROCEED message
		if startupMessage == "PROCEED" {
			return
		}
	}
}

