package tcp

import (
	"cryptocurrency-project/errorchecker"
	"cryptocurrency-project/messages"
	"fmt"
	"math/rand"
	"net"
	"time"
)

// MulticastSendSimulatedDelay waits for a random bounded time and then creates and sends a Message object to all nodes,
// including oneself, via tcp.
func (node Node) MulticastSendSimulatedDelay(message string, nodes []Node) {
	sendTime := time.Now()
	messageWrapper := messages.Message{Message: message, SenderId: node.Id, TimeSent: sendTime}

	rand.Seed(time.Now().UnixNano())
	delay := node.MinDelay + rand.Intn(node.MaxDelay - node.MinDelay)
	fmt.Printf("Delay is %d milliseconds\n", delay)
	time.Sleep(time.Duration(delay) * time.Millisecond)
	SendMessage(nodes, messageWrapper)
}

// SendMessage sends a Message object to all nodes, including oneself via tcp.
func SendMessage(nodes []Node, msg messages.Message) {
	for k, node := range nodes {
		address := node.Ip + ":"
		address += node.Port

		// Attempt to connect to a TCP channel on the localhost IP address.
		// TODO connect to port once instead of constantly reestablishing a connection.
		channel, err := net.Dial("tcp", address)

		//Do not try to encode a message and send it through a tcp channel if it is not active.
		if err != nil {
			fmt.Printf("Node %d could not be reached because it is not listening to its tcp port.\n", k)
			continue
		}

		// Send message through TCP channel.
		err = Encode(channel, msg)
		errorchecker.CheckError(err)
	}
}