package util

import (
	"encoding/hex"
	"strings"

	"github.com/hyperledger/burrow/crypto"
	x "github.com/hyperledger/burrow/encoding/hex"
)

func IrohaAccountIDToAddressHex(accountID string) (string, error) {
	addr := crypto.Keccak256([]byte(accountID))

	return hex.EncodeToString(addr[12:32]), nil
}

func HexToBurrowAddress(h string) (crypto.Address, error) {
	return crypto.AddressFromHexString(x.RemovePrefix(h))
}

func ToEthereumHexString(h string) string {
	return strings.ToLower(x.AddPrefix(h))
}

func ToIrohaHexString(h string) string {
	return strings.ToUpper(x.RemovePrefix(h))
}
