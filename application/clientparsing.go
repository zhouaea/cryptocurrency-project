package application

import (
	"bufio"
	"cryptocurrency-project/tcp"
	"fmt"
	"os"
	"strings"
)

// ReadCommands scans user input from the commandline and converts it into a string.
func ReadCommands(scanner *bufio.Scanner, node tcp.Node, nodes []tcp.Node) string {
	// Loop until user chooses to quit.
	for {
		scanner.Scan()
		text := scanner.Text()

		if text == "quit" {
			os.Exit(0)
		}


		// Convert user input into an array of words for easy parsing.
		words := strings.Fields(text)
		if len(words) < 2 {
			fmt.Println("Please use the format: send <message>")
			continue
		}

		// Parse rest of input into a slice of words that form a message.
		messageSplit := words[1:]

		// Recombine elements of message array to form message string.
		var message string = messageSplit[0]
		for i := 1; i < len(messageSplit); i = i + 1 {
			message = message + " " + messageSplit[i]
		}

		return message
	}
}
