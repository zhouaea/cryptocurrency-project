package tcp

import (
	"cryptocurrency-project/messages"
	"encoding/gob"
	"net"
)

// Encode sends a gob encoded Message object through a connection.
func Encode(conn net.Conn, msg messages.Message) error {
	encoder := gob.NewEncoder(conn)
	err := encoder.Encode(msg)
	return err
}

// Decode receives a gob encoded Message object from a connection and decodes it.
func Decode(conn net.Conn, msg *messages.Message) error {
	decoder := gob.NewDecoder(conn)
	err := decoder.Decode(msg)
	return err
}
