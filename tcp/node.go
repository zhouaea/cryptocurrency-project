// TODO: GET RID OF THIS AFTER USING IT AS SAMPLE CODE
package tcp

import (
	"cryptocurrency-project/message"
	"cryptocurrency-project/telog"
	"fmt"
)

type Node struct {
	Id int
	Ip string
	Port string
	MinDelay int
	MaxDelay int
	UnverifiedTransactions []message.Transaction
	Blockchain telog.Telog
}

// Return a client object from a list of nodes based on their id.
func locateNode(id int, nodes []Node) Node {
	return nodes[id]
}

func (node Node) String() string {
	return fmt.Sprintf("Node ID: %d, Node IP: %s, Node Port: %s Mininimum Delay: %d Maximum Delay %d UnverifiedTransactions %v Blockchain %v", node.Id, node.Ip, node.Port, node.MinDelay, node.MaxDelay, node.UnverifiedTransactions, node.Blockchain)
}