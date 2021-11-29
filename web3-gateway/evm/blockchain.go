package evm

import (
	"encoding/binary"
	"time"

	"github.com/hyperledger/burrow/execution/errors"
)

type blockchain struct {
	blockHeight uint64
	blockTime   time.Time
}

func newBlockchain() *blockchain {
	return &blockchain{}
}

func (b *blockchain) LastBlockHeight() uint64 {
	return b.blockHeight
}

func (b *blockchain) LastBlockTime() time.Time {
	return b.blockTime
}

func (b *blockchain) BlockHash(height uint64) ([]byte, error) {
	if height > b.blockHeight {
		return nil, errors.Codes.InvalidBlockNumber
	}
	bs := make([]byte, 32)
	binary.BigEndian.PutUint64(bs[24:], height)
	return bs, nil
}
