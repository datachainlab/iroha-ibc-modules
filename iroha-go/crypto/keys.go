package crypto

import (
	"crypto"
	"encoding/hex"
	"io"

	"github.com/crpt/go-crpt"
	"github.com/crpt/go-crpt/ed25519"
)

type PrivateKey struct {
	crpt.PrivateKey
}

func GenerateKey(rand io.Reader) (*PublicKey, *PrivateKey, error) {
	c, err := ed25519.New(true, crypto.SHA3_256)
	if err != nil {
		return nil, nil, err
	}

	pubKey, privKey, err := c.GenerateKey(rand)
	if err != nil {
		return nil, nil, err
	}

	return &PublicKey{PublicKey: pubKey}, &PrivateKey{PrivateKey: privKey}, nil
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
