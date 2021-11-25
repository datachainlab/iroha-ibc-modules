package acm

import (
	"errors"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/crypto"
	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
)

type KeyStore interface {
	Set(accountID string, privKey string) error
	SignTransaction(tx *pb.Transaction, accountIDs ...string) ([]*pb.Signature, error)
	SignQuery(query *pb.Query, accountID string) (*pb.Signature, error)
}

var _ KeyStore = (*keyStore)(nil)

type keyStore struct {
	store map[string]string
}

func NewKeyStore() KeyStore {
	return &keyStore{
		store: map[string]string{},
	}
}

var ErrNotExistKey = errors.New("key doesn't exist")
var ErrAlreadyExistKey = errors.New("the key already exists")

func (k keyStore) Set(accountID string, privKey string) error {
	if _, ok := k.store[accountID]; ok {
		return ErrAlreadyExistKey
	} else {
		k.store[accountID] = privKey
		return nil
	}
}

func (k keyStore) SignTransaction(tx *pb.Transaction, accountIDs ...string) ([]*pb.Signature, error) {
	var privKeys []string

	for _, accountID := range accountIDs {
		if privKey, ok := k.store[accountID]; !ok {
			return nil, ErrNotExistKey
		} else {
			privKeys = append(privKeys, privKey)
		}
	}

	sigs, err := crypto.SignTransaction(tx, privKeys...)
	if err != nil {
		return nil, err
	}

	tx.Signatures = sigs

	return sigs, nil
}

func (k keyStore) SignQuery(query *pb.Query, accountID string) (*pb.Signature, error) {
	privKey, ok := k.store[accountID]
	if !ok {
		return nil, ErrNotExistKey
	}

	sig, err := crypto.SignQuery(query, privKey)
	if err != nil {
		return nil, err
	}

	query.Signature = sig

	return sig, nil
}
