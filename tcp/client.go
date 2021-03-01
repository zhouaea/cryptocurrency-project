package tcp

import (
	"MP1/errorchecker"
	"MP1/messages"
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// ExecuteCommands scans user input from the commandline to send a message with bounded delay to another node.
func ExecuteCommands(node Node, nodes []Node) {
	scanner := bufio.NewScanner(os.Stdin)
	// Loop until user chooses to quit.
	for {
		scanner.Scan()
		text := scanner.Text()

		if text == "quit" {
			os.Exit(0)
		}


		// Convert user input into an array of words for easy parsing.
		words := strings.Fields(text)
		if len(words) < 3 {
			fmt.Println("Please use the format: send <destination_id_number> <message>")
			continue
		}

		// Parse destination ID.
		destinationId, err := strconv.Atoi(words[1])
		if destinationId > len(nodes) - 1{
			fmt.Println("Your destination id is out of bounds. Edit the configuration file or use a destination id in the file.")
			continue
		}
		errorchecker.CheckError(err)

		// Parse rest of input into a slice of words that form a message.
		messageSplit := words[2:]

		// Recombine elements of message array to form message string.
		var message string = messageSplit[0]
		for i := 1; i < len(messageSplit); i = i + 1 {
			message = message + " " + messageSplit[i]
		}

		// Send message to destination with unicast.
		go node.UnicastSend(message, destinationId, nodes)
	}
}

// UnicastSend waits for a random bounded time and then creates and sends a Message object to another node via tcp.
func (node Node) UnicastSend(message string, destinationId int, nodes []Node) {
	destinationNode := locateNode(destinationId, nodes)
	sendTime := time.Now()
	messageWrapper := messages.Message{Message: message, SenderId: node.Id, TimeSent: sendTime}

	delay := node.MinDelay + rand.Intn(node.MaxDelay - node.MinDelay)
	fmt.Printf("Delay is %d milliseconds\n", delay)
	time.Sleep(time.Duration(delay) * time.Millisecond)
	SendMessage(destinationNode, messageWrapper)
}

// SendMessage sends a Message object to another node via tcp.
func SendMessage(destinationNode Node, msg messages.Message) {
	address := "127.0.0.1:"
	address += destinationNode.Port

	// Attempt to connect to a TCP channel on the localhost IP address.
	channel, err := net.Dial("tcp", address)
	errorchecker.CheckError(err)

	// Send message through TCP channel.
	messages.Encode(channel, msg)
}