package messages

import (
	"time"
)

type Message struct {
	Message string
	SenderId int
	TimeSent time.Time
}

type Transaction struct {
	Sender string
	Receiver string
	Value int
	Change int
	Signature string
}