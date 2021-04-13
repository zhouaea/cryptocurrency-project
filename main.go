package main

import (
	"./initialization"
	"./tcp"
	"fmt"
	"./telog"
)

func main() {
	// Read information from command line and configuration file to obtain the node the process should act as
	// and a list of valid nodes to send messages to.
	node, nodes := initialization.InitializeNode()

	// Set up TCP listening for process, continuously listening for messages and printing them out.
	go node.UnicastReceive()

	// Continuously scan user input for instructions on which node to send a message to.
	tcp.ExecuteCommands(node, nodes)

	log := telog.Telog{}
	log.Init()

	log.AddBlock("Goofy mints 5 dollars")
	log.AddBlock("Goofy paid Alice 5 dollars")
	log.AddBlock("Alice paid Bob 5 dollars")

	fmt.Println("# Blocks in Log:", log.GetNumBlocks())

	fmt.Println("Log is valid:", log.Check())

	log.Attack(1)

	fmt.Println("After attack, Log is valid:", log.Check())
}


