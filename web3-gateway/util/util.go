package util

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/hyperledger/burrow/crypto"
)

func IrohaAccountIDToAddressHex(accountID string) string {
	addr := crypto.Keccak256([]byte(accountID))

	return hex.EncodeToString(addr[12:32])
}

func ToEthereumHexString(h string) string {
	return strings.ToLower(AddHexPrefix(h))
}

func ToIrohaHexString(h string) string {
	return strings.ToUpper(RemoveHexPrefix(h))
}

func AddHexPrefix(h string) string {
	if has0xPrefix(h) {
		return h
	}

	return fmt.Sprintf("0x%s", h)
}

func RemoveHexPrefix(h string) string {
	if !has0xPrefix(h) {
		return h
	}

	return strings.TrimPrefix(h, "0x")
}

func has0xPrefix(s string) bool {
	return len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X')
}
