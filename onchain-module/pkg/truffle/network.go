package truffle

// generated by "schematyper -o network.go --package truffle ../../node_modules/@truffle/contract-schema/spec/network-object.spec.json" -- DO NOT EDIT

type Address string

type NetworkObject struct {
	Address         Address                `json:"address,omitempty"`
	Db              map[string]interface{} `json:"db,omitempty"`
	Events          map[string]interface{} `json:"events,omitempty"`
	Links           map[string]interface{} `json:"links,omitempty"`
	TransactionHash TransactionHash        `json:"transactionHash,omitempty"`
}

type TransactionHash string
