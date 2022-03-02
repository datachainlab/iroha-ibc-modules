package main

import (
	"log"

	"github.com/datachainlab/ibc-ethmultisig-client/modules/relay/ethmultisig"
	"github.com/hyperledger-labs/yui-relayer/cmd"

	iroha "github.com/datachainlab/iroha-ibc-modules/relayer/chains/iroha/module"
)

func main() {
	err := cmd.Execute(
		iroha.Module{},
		ethmultisig.Module{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
