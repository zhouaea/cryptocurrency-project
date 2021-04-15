package main

import (
	"./telog"
	"fmt"
)

func main() {
	log := telog.Telog{}
	log.Init()

	//TODO Add a block using a transaction object.
	//log.AddBlock("Goofy mints 5 dollars")
	//log.AddBlock("Goofy paid Alice 5 dollars")
	//log.AddBlock("Alice paid Bob 5 dollars")

	fmt.Println("# Blocks in Log:", log.GetNumBlocks())

	fmt.Println("Log is valid:", log.Check())

	log.Attack(1)

	fmt.Println("After attack, Log is valid:", log.Check())
}


