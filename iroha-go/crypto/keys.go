package crypto

import (
	"encoding/hex"

	"github.com/crpt/go-crpt"
)

type PrivateKey struct {
	crpt.PrivateKey
}

func (sk *PrivateKey) Address() []byte {
	return sk.Bytes()[:32]
}

func (sk *PrivateKey) Hex() string {
	return hex.EncodeToString(sk.Address())
}

type PublicKey struct {
	crpt.PublicKey
}

func (pk *PublicKey) Hex() string {
	return hex.EncodeToString(pk.Address())
}
