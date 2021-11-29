package config

import (
	"time"
)

type Config struct {
	Iroha    Iroha          `json:"iroha" yaml:"iroha"`
	Gateway  Gateway        `json:"gateway" yaml:"gateway"`
	Accounts []IrohaAccount `json:"accounts" yaml:"accounts"`
	EVM      EVM            `json:"evm" yaml:"evm"`
}

type Iroha struct {
	Api struct {
		Host           string        `json:"host" yaml:"host"`
		Port           int           `json:"port" yaml:"port"`
		CommandTimeout time.Duration `json:"commandTimeout" yaml:"commandTimeout"`
		QueryTimeout   time.Duration `json:"queryTimeout" yaml:"queryTimeout"`
	} `json:"api" yaml:"api"`

	Database struct {
		Postgres struct {
			Host     string `json:"host" yaml:"host"`
			Port     int    `json:"port" yaml:"port"`
			User     string `json:"user" yaml:"user"`
			Password string `json:"password" yaml:"password"`
			Database string `json:"database" yaml:"database"`
		} `json:"postgres" yaml:"postgres"`
	} `json:"database" yaml:"database"`
}

type Gateway struct {
	Rpc struct {
		Host string `json:"host" yaml:"host"`
		Port int    `json:"port" yaml:"port"`
	} `json:"rpc" yaml:"rpc"`
}

type IrohaAccount struct {
	ID         string `json:"id" yaml:"id"`
	PrivateKey string `json:"privateKey" yaml:"privateKey"`
}

type EVM struct {
	Querier string `json:"querier" yaml:"querier"`
}
