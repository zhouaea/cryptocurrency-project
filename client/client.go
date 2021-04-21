package client

import (
	"../esign"
	"crypto/ed25519"
	"cryptocurrency-project/errorchecker"
	"cryptocurrency-project/ipaddresses"
	"cryptocurrency-project/message"
	"cryptocurrency-project/tcp"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
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

// SendStartup sends a start up message to the controller.
func SendStartup(clientID int) net.Conn {
	message := message.Message{"READY", "Client " + string(clientID)}
	controllerChannel, err := net.Dial("tcp", ipaddresses.GetController())
	if err != nil {
		fmt.Println("Controller has not started up yet. Try again after booting up controller.")
		os.Exit(1)
	}
	tcp.Encode(controllerChannel, message)

	return controllerChannel
}

// WaitForController waits for a PROCEED message (string, not an object) from the controller that all processes have
// started up.
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

// SendRequest() periodically multicasts transactions to miners.
func SendRequest(delayInterval int) {
	// Create and record connections to each miner.
	minerChannels := make(map[string]net.Conn)

	minerAddresses := ipaddresses.GetMiners()
	for i := 0; i < len(minerAddresses); i++ {
		minerChannel, err := net.Dial("tcp", minerAddresses[i])
		errorchecker.CheckError(err)
		minerChannels["Miner " + string(i)] = minerChannel
	}

	// TODO Create an iterable data structure of transactions to periodically send out transactions, one for each client ID
	for transaction := range transactions {
		go MulticastSendToMinersSimulatedDelay(transaction, minerChannels)
		// Wait a set amount of time before sending another transaction to the miners.
		time.Sleep(time.Duration(delayInterval) * time.Millisecond)
	}
}

// MulticastSendSimulatedDelay waits for a random bounded time to simulate message delay and then sends a transaction
// object to specified destinations via tcp.
func MulticastSendToMinersSimulatedDelay(clientTransaction Transaction, minerChannels map[string]net.Conn) {
	// Delay sending message for a random bounded time in milliseconds
	minDelay := 0
	maxDelay := 10000
	rand.Seed(time.Now().UnixNano())
	delay := minDelay + rand.Intn(maxDelay - minDelay)
	fmt.Printf("Delay is %d milliseconds\n", delay)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	// Send message through TCP channel.
	for _, minerChannel := range minerChannels {
		err := tcp.Encode(minerChannel, clientTransaction)
		errorchecker.CheckError(err)
	}
}