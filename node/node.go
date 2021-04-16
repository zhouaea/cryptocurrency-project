package node

import (
	"../esign"
	"crypto/ed25519"
)

type node struct {
	Name       string
	Sk ed25519.PrivateKey
	Pk ed25519.PublicKey
}

type Nodes struct {
	all []node
}

// Creates a new node
func (n *Nodes) NewNode(name string) node {
	pk, sk := esign.GenerateKeyPair()

	new := node{
		Name: name,
		Sk: sk,
		Pk: pk,
	}

	n.all = append(n.all, new)
	return new
}
