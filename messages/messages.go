package messages

import (
	"encoding/gob"
	"net"
	"time"
)

type Message struct {
	Message string
	SenderId int
	TimeSent time.Time
}
// Encode sends a gob encoded Message object through a connection.
func Encode(conn net.Conn, msg Message) {
	encoder := gob.NewEncoder(conn)
	encoder.Encode(msg)
}

// Decode receives a gob encoded Message object from a connection and decodes it.
func Decode(conn net.Conn, msg *Message) {
	decoder := gob.NewDecoder(conn)
	decoder.Decode(msg)
}
