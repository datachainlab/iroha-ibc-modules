package iroha

import (
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/api"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db"
)

type Client struct {
	api.ApiClient
	db.DBClient
}

func New(apiClient api.ApiClient, dbClient db.DBClient) *Client {
	return &Client{
		ApiClient: apiClient,
		DBClient:  dbClient,
	}
}
