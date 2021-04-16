/*

	The tx package implements transactions for cryptocurrency.

*/
package tx

import (
	"../esign"
	"crypto/ed25519"
	"fmt"
	"time"
)

var pk1, sk1 = esign.GenerateKeyPair()

var CoinMakerPk = pk1
var CoinMakerSk = sk1

type tx struct {
	timestamp int64
	senderName string
	receiverName string
	senderPk *ed25519.PublicKey
	receiverPk *ed25519.PublicKey
	amountInput int
	amountToSelf int
	amountToReceiver int
	signature []byte
}

type TxArray struct {
	txs []tx
}

// Add a new tx to txArray
func (txa *TxArray) AppendNewTx(
	senderName string,
	receiverName string,
	senderPk *ed25519.PublicKey,
	receiverPk *ed25519.PublicKey,
	senderSk *ed25519.PrivateKey,
	value int,
) {
	var amountInput int
	var amountToReceiver int
	var amountToSelf int

	// If senderPk is CoinMakerPk, then create coin
	if senderPk.Equal(CoinMakerPk) {
		amountInput = value
		amountToSelf = 0
		amountToReceiver = value
	} else {
		coinOwned := 0
		mostRecentExpense := tx{}

		// If senderPk is a node, loop through the transaction array
		for i := len(txa.txs)-1; i >= 0; i-- {
			tx := txa.txs[i]

			// Verify signature of transaction
			if !esign.VerifyTx(*tx.senderPk, tx.timestamp, tx.amountInput, tx.amountToSelf, tx.amountToReceiver, tx.signature) {
				fmt.Println("Tx failed. Failed to verify past transactions.")
				return
			}

			if mostRecentExpense.senderPk != nil {
				// If pk of most recent expense = most recent income
				if (*tx.receiverPk).Equal(*mostRecentExpense.senderPk) {
					coinOwned += mostRecentExpense.amountToSelf
					break
				}
			}

			// Sum up unspent coins
			if (*tx.receiverPk).Equal(*senderPk) {
				coinOwned += tx.amountToReceiver
			}

			// Store the most recent time that the node spent coins
			if (*tx.senderPk).Equal(*senderPk) {
				mostRecentExpense = tx
			}
		}

		// Check if you have enough coin
		if coinOwned < value {
			fmt.Println("Tx failed. Not enough coins")
			return
		}

		amountInput = coinOwned
		amountToReceiver = value
		amountToSelf = coinOwned - value

	}

	timestamp := time.Now().UnixNano()
	signature := esign.SignTx(*senderSk, timestamp, amountInput, amountToSelf, amountToReceiver)

	newTx := tx{
		timestamp: timestamp,
		senderName: senderName,
		receiverName: receiverName,
		senderPk: senderPk,
		receiverPk: receiverPk,
		amountInput: amountInput,
		amountToSelf: amountToSelf,
		amountToReceiver: amountToReceiver,
		signature: signature,
	}

	txa.txs = append(txa.txs, newTx)
}

func (txa *TxArray) PrintTxHistory() {
	for i, tx := range txa.txs {
		fmt.Println("Tx Index:", i)
		fmt.Println(tx.senderName, "--->", tx.receiverName)
		fmt.Println("INPUT:", tx.amountInput)
		fmt.Println("OUTPUT: to receiver:", tx.amountToReceiver, "& to self:", tx.amountToSelf)
		fmt.Println()
	}
}