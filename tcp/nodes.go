package tcp

import (
	"fmt"
)

type Node struct {
	Id int
	Ip string
	Port string
	MinDelay int
	MaxDelay int
}

// Return a node object from a list of nodes based on their id.
func locateNode(id int, nodes []Node) Node {
	return nodes[id]
}

func (node Node) String() string {
	return fmt.Sprintf(
		"Node ID: %d, Node IP: %s, Node Port: %s Mininimum Delay: %d Maximum Delay %d",
		node.Id,
		node.Ip,
		node.Port,
		node.MinDelay,
		node.MaxDelay,
	)
}