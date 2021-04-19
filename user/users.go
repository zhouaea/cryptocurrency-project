package user

import (
	"../esign"
	"crypto/ed25519"
	"cryptocurrency-project/ipaddresses"
	"cryptocurrency-project/message"
	"cryptocurrency-project/tcp"
)

type user struct {
	Name       string
	Sk ed25519.PrivateKey
	Pk ed25519.PublicKey
}

type Users struct {
	all []user
}

// Creates a new user
func (n *Users) NewUser(name string) user {
	pk, sk := esign.GenerateKeyPair()

	new := user{
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
