package miner

import (
	"cryptocurrency-project/errorchecker"
	"cryptocurrency-project/ipaddresses"
	"cryptocurrency-project/message"
	"cryptocurrency-project/tcp"
	"cryptocurrency-project/telog"
	"cryptocurrency-project/tx"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

type MinerStorage struct {
	Blockchain telog.Telog
	UnverifiedTransactions tx.TxArray
}

// SendStartup sends a start up message to the controller
func SendStartup(minerID int) net.Conn {
	message := message.Message{"READY", "Miner " + string(minerID)}
	controllerChannel, err := net.Dial("tcp", ipaddresses.GetController())
	if err != nil {
		fmt.Println("Controller has not started up yet. Try again after booting up controller.")
		os.Exit(1)
	}
	tcp.Encode(controllerChannel, message)

	return controllerChannel
}

// ConnectToClients stores communication channels to each client to listen to later.
func ConnectToClients() []net.Conn {
	clientAddresses := ipaddresses.GetClients()
	clientChannels := make([]net.Conn, 0, len(clientAddresses))

	numberOfClients := len(clientAddresses)
	for i := 0; i < numberOfClients; i++ {
		clientChannel, err := net.Dial("tcp", clientAddresses[i])
		errorchecker.CheckError(err)
		clientChannels[i] = clientChannel
	}

	return clientChannels
}

// ConnectToOtherMiners will connect to other miners in a way that each pair of miner combinations requires only one
// tcp channel.
func ConnectToOtherMiners(minerId int, minerChannels []net.Conn) {
	minerAddresses := ipaddresses.GetMiners()
	
	// Pattern: miner 0 connects to all miners except itself, miner 1 connects to all miners except miner 0 and itself, etc.
	for i := minerId + 1; i < len(minerAddresses); i++ {
		minerChannel, err := net.Dial("tcp", minerAddresses[i])
		errorchecker.CheckError(err)
		// Appending should never force the allocation of a new memory address since the capacity of the slice is equal
		// to the number of miners.
		minerChannels = append(minerChannels, minerChannel)
	}
}

func ListenToOtherMiners(listener net.Listener, minerChannels []net.Conn) {
	// Wait for miner connections until all miners except ourselves are connected to us.
	for len(minerChannels) < len(ipaddresses.GetMiners()) - 1 {
		// Wait for a connection from a miner to our TCP port and then set up a TCP channel with them.
		minerChannel, err := listener.Accept()
		errorchecker.CheckError(err)

		minerChannels = append(minerChannels, minerChannel)
	}
}


// MiningProtocol performs 3 actions:
// 1. Upon reception of ClientTransactionRequest objects, add them to the transaction array.
// 2. Upon reception of MINING SUCCESSFUL message from controller, propose block with transactions in transaction array.
// 3. Upon reception of PROPOSAL message, verify that the transactions are valid with correct signatures. If so,
//    telog.AddBlock(). Tell the controller that a PROPOSAL message was verified or denied with PROPOSAL RECEIVED.
func MiningProtocol(controllerChannel net.Conn, clientChannels []net.Conn, otherMinerChannels []net.Conn, listener net.Listener, minerStoragePointer *MinerStorage) {
	go listenToClients(clientChannels, minerStoragePointer)
	go listenToController(controllerChannel, otherMinerChannels, minerStoragePointer)
	go listenToMiners(otherMinerChannels, minerStoragePointer)
}

// Upon reception of ClientTransactionRequest objects, add them to the transaction array.
func listenToClients(clientChannels []net.Conn, minerStoragePointer *MinerStorage) {
	for _, clientChannel := range clientChannels{
		go listenToClient(clientChannel, minerStoragePointer)
	}
}

func listenToClient(clientChannel net.Conn, minerStoragePointer *MinerStorage) {
	// TODO Use real transaction struct
	var transaction string
	tcp.Decode(clientChannel, transaction)
	// minerStoragePointer.UnverifiedTransactions.AppendNewTx(transaction)
}

// Upon reception of MINING SUCCESSFUL message from controller, propose block with transactions in transaction array.
func listenToController(controllerChannel net.Conn, otherMinerChannels []net.Conn, minerStoragePointer *MinerStorage) {
	for {
		var message string
		tcp.Decode(controllerChannel, message)

		if message == "MINING SUCCESSFUL" {
			proposeSimulatedDelay(otherMinerChannels, minerStoragePointer.UnverifiedTransactions.Dequeue())
		}
	}
}

// Propose a block to all nodes.
func proposeSimulatedDelay(otherMinerChannels []net.Conn, transaction string) {
	// Delay sending message for a random bounded time in milliseconds
	minDelay := 0
	maxDelay := 10000
	rand.Seed(time.Now().UnixNano())
	delay := minDelay + rand.Intn(maxDelay - minDelay)
	fmt.Printf("Delay is %d milliseconds\n", delay)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	// Send proposal through TCP channel.
	for _, minerChannel := range otherMinerChannels {
		err := tcp.Encode(minerChannel, transaction)
		errorchecker.CheckError(err)
	}
}

// Upon reception of PROPOSAL message, verify that the transactions are valid with correct signatures. If so,
// telog.AddBlock(). Tell the controller that a PROPOSAL message was verified or denied with PROPOSAL RECEIVED.
func listenToMiners(otherMinerChannels []net.Conn, minerStoragePointer *MinerStorage) {
	for _, minerChannel := range otherMinerChannels{
		go listenToMiner(minerChannel, minerStoragePointer)
	}
}

func listenToMiner(minerChannel net.Conn, minerStoragePointer *MinerStorage) {
	// TODO Use real transaction struct
	var transaction string
	tcp.Decode(minerChannel, transaction)
	verify(transaction)
}

// Ensure transaction is valid before adding to blockchain.
func verify(transaction string) {

	// minerStoragePointer.Blockchain.AddBlock(transaction)
}
