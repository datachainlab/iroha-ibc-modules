package crypto

import (
	"bytes"
	"crypto"
	"encoding/hex"

	"github.com/crpt/go-crpt"
	"github.com/crpt/go-crpt/ed25519"
	"github.com/golang/protobuf/proto"
	"golang.org/x/crypto/sha3"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
)

func SignTransaction(tx *protocol.Transaction, privKeys ...string) ([]*protocol.Signature, error) {
	var sigs []*protocol.Signature

	for _, privKey := range privKeys {
		sig, pubKey, err := signature(tx.Payload, privKey)
		if err != nil {
			return nil, err
		}

		sigs = append(sigs, &protocol.Signature{
			Signature: hex.EncodeToString(sig),
			PublicKey: hex.EncodeToString(pubKey.Address()),
		})
	}

	return sigs, nil
}

func SignQuery(query *protocol.Query, privKey string) (*protocol.Signature, error) {
	sig, pubKey, err := signature(query.Payload, privKey)
	if err != nil {
		return nil, err
	}

	return &protocol.Signature{
		Signature: hex.EncodeToString(sig),
		PublicKey: hex.EncodeToString(pubKey.Address()),
	}, nil
}

func signature(message proto.Message, privKeyHex string) ([]byte, crpt.PublicKey, error) {
	c, err := ed25519.New(true, crypto.SHA3_256)
	if err != nil {
		return nil, nil, err
	}

	seed, err := hex.DecodeString(privKeyHex)
	if err != nil {
		return nil, nil, err
	}

	pubKey, privKey, err := c.GenerateKey(bytes.NewReader(seed))
	if err != nil {
		return nil, nil, err
	}

	digest, err := Hash(message)
	if err != nil {
		return nil, nil, err
	}

	sig, err := c.Sign(privKey, digest, nil, crpt.NotHashed, nil)
	if err != nil {
		return nil, nil, err
	}

	return sig, pubKey, nil
}

func Hash(message proto.Message) ([]byte, error) {
	bz, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}

	digest := sha3.Sum256(bz)
	return digest[:], nil
}
