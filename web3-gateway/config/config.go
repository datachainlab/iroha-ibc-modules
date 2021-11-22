package config

type Config struct {
	Iroha   Iroha   `json:"iroha" yaml:"iroha"`
	Gateway Gateway `json:"gateway" yaml:"gateway"`
}

type Iroha struct {
	Api struct {
		Host string `json:"host" yaml:"host"`
		Port int    `json:"port" yaml:"port"`
	} `json:"api" yaml:"api"`

	Postgres struct {
		Host     string `json:"host" yaml:"host"`
		Port     int    `json:"port" yaml:"port"`
		User     string `json:"user" yaml:"user"`
		Password string `json:"password" yaml:"password"`
		Database string `json:"database" yaml:"database"`
	} `json:"postgres" yaml:"postgres"`
}

type Gateway struct {
	Rpc struct {
		Host string `json:"host" yaml:"host"`
		Port int    `json:"port" yaml:"port"`
	} `json:"rpc" yaml:"rpc"`
	Accounts []IrohaAccount `json:"accounts" yaml:"accounts"`
}

type IrohaAccount struct {
	ID         string `json:"id" yaml:"id"`
	PrivateKey string `json:"privateKey" yaml:"privateKey"`
}
