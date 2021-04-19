package main

import (
	"cryptocurrency-project/errorchecker"
	"cryptocurrency-project/ipaddresses"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read command line for user id to act as.
	if len(os.Args) != 2 {
		fmt.Println("Format: go run minerprocess.go <client_index> ")
	}
	nodeIndex, err := strconv.Atoi(os.Args[1])
	errorchecker.CheckError(err)

	// Connect to specified tcp port of miner.
	minerAddress := ipaddresses.GetMiners()
	port := strings.Split(minerAddress[nodeIndex], ":")[1]
	listener, err := net.Listen("tcp", port)

	// Send startup message to controller.
	SendStartup(sender: minerAddress)

	for {
		// Upon reception of ClientTransactionRequest objects, add them to the transaction array.
		if (message == "REQUEST") {
			ReceiveClientTransactions()
		}

		// Upon reception of CHOSEN message, propose block with transactions in transaction array.
		if (message == "CHOSEN") {
			Propose()
		}

		// Upon receival of propose message, verify that the transactions are valid with correct signatures. If so, telog.AddBlock().
		// Tell the controller that a propose message was received with PROPOSE RECEIVED.
		if (message == "PROPOSE") {
			VerifyPropose()
		}
	}

}
