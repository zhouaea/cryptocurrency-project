package client

import (
	"../esign"
	"crypto/ed25519"
	"cryptocurrency-project/ipaddresses"
	"cryptocurrency-project/message"
	"cryptocurrency-project/tcp"
)

type client struct {
	Name       string
	Sk ed25519.PrivateKey
	Pk ed25519.PublicKey
}

type Clients struct {
	all []client
}

// Creates a new client
func (n *Clients) Newclient(name string) client {
	pk, sk := esign.GenerateKeyPair()

	new := client{
		Name: name,
		Sk: sk,
		Pk: pk,
	}

	n.all = append(n.all, new)
	return new
}

func SendStartup(clientAddress string) {
	message := message.InitialConnection{"STARTED UP", clientAddress}
	ipaddresses.GetController()
	tcp.Encode()
}
