package main

import (
	"./node"
	"./tx"
)

func main() {
	//log := telog.Telog{}
	//log.Init()
	//
	//log.AddBlock("Goofy mints 5 dollars")
	//log.AddBlock("Goofy paid Alice 5 dollars")
	//log.AddBlock("Alice paid Bob 5 dollars")
	//
	//fmt.Println("# Blocks in Log:", log.GetNumBlocks())
	//
	//fmt.Println("Log is valid:", log.Check())
	//
	//log.Attack(1)
	//
	//fmt.Println("After attack, Log is valid:", log.Check())

	transactions := tx.TxArray{}

	nodes := node.Nodes{}
	alice := nodes.NewNode("Alice")
	bob := nodes.NewNode("Bob")

	transactions.AppendNewTx("CoinMaker", "Bob", &tx.CoinMakerPk, &bob.Pk, &tx.CoinMakerSk, 200)
	transactions.AppendNewTx("CoinMaker", "Alice", &tx.CoinMakerPk, &alice.Pk, &tx.CoinMakerSk, 100)
	transactions.AppendNewTx("CoinMaker", "Bob", &tx.CoinMakerPk, &bob.Pk, &tx.CoinMakerSk, 200)
	transactions.AppendNewTx("Alice", "Bob", &alice.Pk, &bob.Pk, &alice.Sk, 50)
	transactions.AppendNewTx("Bob", "Alice", &bob.Pk, &alice.Pk, &bob.Sk, 150)

	transactions.PrintTxHistory()

}


