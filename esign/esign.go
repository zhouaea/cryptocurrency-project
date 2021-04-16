package esign

import (
	"bytes"
	"crypto/ed25519"
	"crypto/sha256"
	"fmt"
	"strconv"
)

// Generates pk and sk
func GenerateKeyPair() (ed25519.PublicKey, ed25519.PrivateKey) {
	pk, sk, err := ed25519.GenerateKey(nil)
	if err != nil {
		fmt.Println("Failed to generate pk and sk")
	}
	return pk, sk
}

func hashSha256(b []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(b))
}

func hashMsg(timestamp int64, amountInput int, amountToSelf int, amountToReceiver int) []byte {
	t := []byte(strconv.FormatInt(timestamp, 10))
	a1 := []byte(strconv.Itoa(amountInput))
	a2 := []byte(strconv.Itoa(amountToSelf))
	a3 := []byte(strconv.Itoa(amountToReceiver))
	msg := bytes.Join([][]byte{t, a1, a2, a3}, []byte{})
	hashedMsg := []byte(hashSha256(msg))
	return hashedMsg
}

// Signs a transaction
func SignTx(sk ed25519.PrivateKey, timestamp int64, amountInput int, amountToSelf int, amountToReceiver int) []byte {
	hashedMsg := hashMsg(timestamp, amountInput, amountToSelf, amountToReceiver)
	return ed25519.Sign(sk, hashedMsg)
}

// Verifies a transaction
func VerifyTx(pk ed25519.PublicKey, timestamp int64, amountInput int, amountToSelf int, amountToReceiver int, signature []byte) bool {
	hashedMsg := hashMsg(timestamp, amountInput, amountToSelf, amountToReceiver)
	flag := ed25519.Verify(pk, hashedMsg, signature)
	return flag
}