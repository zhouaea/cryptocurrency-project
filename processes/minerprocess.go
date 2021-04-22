package main

import (
	"cryptocurrency-project/errorchecker"
	"cryptocurrency-project/ipaddresses"
	"cryptocurrency-project/miner"
	"cryptocurrency-project/telog"
	"cryptocurrency-project/tx"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read command line for client id to act as.
	if len(os.Args) != 2 {
		fmt.Println("Format: go run minerprocess.go <client_index> ")
	}
	minerIndex, err := strconv.Atoi(os.Args[1])
	errorchecker.CheckError(err)

	// Connect to specified tcp port of miner.
	minerAddresses := ipaddresses.GetMiners()
	port := strings.Split(minerAddresses[minerIndex], ":")[1]
	listener, err := net.Listen("tcp", port)

	// Send startup message to controller.
	controllerChannel := miner.SendStartup(minerIndex)

	// Initialize connection to clients.
	clientChannels := miner.ConnectToClients()

	// Initialize connection to other miners.
	minerChannels := make ([]net.Conn, 0, len(ipaddresses.GetMiners()))
	go miner.ConnectToOtherMiners(minerIndex, minerChannels)
	go miner.ListenToOtherMiners(listener, minerChannels)


	// Create a miner object to represent the data stored in a miner.
	minerStorage := miner.MinerStorage{Blockchain: telog.Telog{}, UnverifiedTransactions: tx.TxArray{}}

	// Upon reception of ClientTransactionRequest objects, add them to the transaction array.
	// Upon reception of MINING SUCCESSFUL message, propose block with transactions in transaction array.
	// Upon receival of propose message, verify that the transactions are valid with correct signatures. If so, telog.AddBlock().
	//     Tell the controller that a propose message was received with PROPOSE RECEIVED.
	miner.MiningProtocol(controllerChannel, clientChannels, minerChannels, listener, &minerStorage)


}
