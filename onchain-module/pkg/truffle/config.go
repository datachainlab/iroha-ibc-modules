package truffle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
)

//go:generate schematyper -o network.go --package truffle ../../node_modules/@truffle/contract-schema/spec/network-object.spec.json

type Config struct {
	Networks map[string]NetworkObject `json:"networks"`
}

func UnmarshallConfig(networkID, buildPath, fName string) NetworkObject {
	raw, err := ioutil.ReadFile(path.Clean(path.Join(buildPath, fName)))
	if err != nil {
		panic(err)
	}
	tc := &Config{}
	err = json.Unmarshal(raw, &tc)
	if err != nil {
		panic(err)
	}
	n, ok := tc.Networks[networkID]
	if !ok {
		panic(fmt.Errorf("the contract is not deployed to network ID=%s FileName=%s", networkID, fName))
	}

	if len(n.Address) == 0 {
		panic(fmt.Errorf("contract address is empty network ID=%s FileName=%s", networkID, fName))
	}

	return n
}
