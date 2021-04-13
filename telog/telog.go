/*

	The telog package is an implementation of the Tamper-evident log.
	It supports three major APIs:
		1. Init(): Initialize a tamper-evident log.
		2. AddBlock(data): Add block to the end of the chain.
		3. Check(): Iterate through a log and checks whether there is some block that has been tampered.
	And two APIs for testing purposes:
		1. GetNumBlocks(): Return the number of blocks in the log.
		2. Attack(idx): Modify data in block at position idx.

*/
package telog

import (
	"crypto/sha256"
	"fmt"
)

type hashPointer struct {
	pointer *block
	hash string
}

type block struct {
	hashPointer hashPointer
	data string
}

type Telog struct {
	head hashPointer
	count int
}

// Init initializes an empty head used for the tamper evident log with SHA-256.
func (t *Telog) Init() {
	t.head = hashPointer{}
	t.count = 0
}

func (t *Telog) GetNumBlocks() int {
	return t.count
}

// hashSha256 returns the hash digest of a block.
func (t *Telog) hashSha256(block block) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%v", block))))
}

// AddBlock adds a block with data to the right end of a tamper evident log, where the left end of the log is the first
// data block added and the right end of the log is the last data block added.
func (t *Telog) AddBlock(data string) {
	// Create a new block.
	newBlock := block{
		// Use the old hash pointer of the head to connect the new block to the right-most block
		hashPointer: t.head,
		data: data,
	}
	
	// Hash the new block.
	newBlockHash := t.hashSha256(newBlock)

	// Connect the head to the new block with a hash pointer.
	t.head = hashPointer{
		pointer: &newBlock,
		hash: newBlockHash,
	}

	t.count += 1
}

// Check determines if the log has been tampered with, returning true if the log is valid and false if the log was
// tampered with.
func (t *Telog) Check() bool {
	currentHashPointer := t.head
	emptyHashPointer := hashPointer{}

	// Execute as long as there is not a non-empty hash pointer.
	for currentHashPointer != emptyHashPointer {
		// Access the block pointed to by the hash pointer
		currentBlock := *currentHashPointer.pointer

		// Rehash the block to check if it was tampered with.
		currentBlockHash := t.hashSha256(currentBlock)
		if currentBlockHash	!= currentHashPointer.hash {
			return false
		}

		// Iterate to next pointer
		currentHashPointer = currentBlock.hashPointer
	}
	return true
}

// Attack modifies the block at position idx.
// Position 0 is the genesis block.
func (t *Telog) Attack(idx int) {
	// idx starts counting from the genesis block
	// reverseIdx starts counting from the head
	reverseIdx := t.count - idx
	if reverseIdx <= 0 {
		fmt.Println("The block to attack does not exist.")
		return
	}

	// Find the pointer to the block that is being attacked
	currentHashPointer := t.head
	for i := 1; i < reverseIdx; i++{
		currentHashPointer = (*currentHashPointer.pointer).hashPointer
	}

	// Change data in that block to "Attacked"
	(*currentHashPointer.pointer).data = "Attacked"

	fmt.Println("ATTACK: Block", idx)
}