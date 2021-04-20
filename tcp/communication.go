package tcp

import (
	"encoding/gob"
	"net"
)

// Encode sends a gob encoded message through a connection.
func Encode(conn net.Conn, msg interface{}) error {
	encoder := gob.NewEncoder(conn)
	err := encoder.Encode(msg)
	return err
}

// Decode receives a gob encoded message from a connection and decodes it.
func Decode(conn net.Conn, msg interface{}) error {
	decoder := gob.NewDecoder(conn)
	err := decoder.Decode(msg)
	return err
}
