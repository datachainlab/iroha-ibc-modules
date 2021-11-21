package query

import (
	"context"
	"errors"
	"time"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/crypto"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
)

type QueryClient interface {
	GetAccount(accountID string, opts ...PayLoadMetaOption) *protocol.Query
	GetSignatories(accountID string, opts ...PayLoadMetaOption) *protocol.Query
	GetAccountTransactions(accountID string, txPaginationMeta *protocol.TxPaginationMeta, opts ...PayLoadMetaOption) *protocol.Query
	GetAccountAssetTransactions(accountID string, assetId string, txPaginationMeta *protocol.TxPaginationMeta, opts ...PayLoadMetaOption) *protocol.Query
	GetTransactions(txHashes []string, opts ...PayLoadMetaOption) *protocol.Query
	GetAccountAsset(accountID string, paginationMeta *protocol.AssetPaginationMeta, opts ...PayLoadMetaOption) *protocol.Query
	GetAccountDetail(accountID string, key string, writer string, paginationMeta *protocol.AccountDetailPaginationMeta, opts ...PayLoadMetaOption) *protocol.Query
	GetRoles(opts ...PayLoadMetaOption) *protocol.Query
	GetRolePermissions(roleID string, opts ...PayLoadMetaOption) *protocol.Query
	GetAssetInfo(assetID string, opts ...PayLoadMetaOption) *protocol.Query
	GetPendingTransactions(txPaginationMeta *protocol.TxPaginationMeta, opts ...PayLoadMetaOption) *protocol.Query
	GetBlock(height uint64, opts ...PayLoadMetaOption) *protocol.Query
	GetPeers(opts ...PayLoadMetaOption) *protocol.Query
	GetEngineReceipts(txHash string, opts ...PayLoadMetaOption) *protocol.Query

	SendQuery(ctx context.Context, query *protocol.Query, privKeyHex string) (*protocol.QueryResponse, error)
}

var _ QueryClient = (*queryClient)(nil)

type queryClient struct {
	client protocol.QueryServiceV1Client

	Timeout time.Duration
}

func New(client protocol.QueryServiceV1Client, timeout time.Duration) QueryClient {
	return &queryClient{
		client:  client,
		Timeout: timeout,
	}
}

type PayLoadMetaOption func(*protocol.QueryPayloadMeta)

func CreatorAccountId(id string) PayLoadMetaOption {
	return func(meta *protocol.QueryPayloadMeta) {
		meta.CreatorAccountId = id
	}
}

func QueryCounter(counter uint64) PayLoadMetaOption {
	return func(meta *protocol.QueryPayloadMeta) {
		meta.QueryCounter = counter
	}
}

func CreatedTime(t uint64) PayLoadMetaOption {
	return func(meta *protocol.QueryPayloadMeta) {
		meta.CreatedTime = t
	}
}

func (c *queryClient) payloadMeta(opts ...PayLoadMetaOption) *protocol.QueryPayloadMeta {
	meta := &protocol.QueryPayloadMeta{
		QueryCounter: 1,
		CreatedTime:  uint64(time.Now().UnixNano() / int64(time.Millisecond)),
	}

	for _, opt := range opts {
		opt(meta)
	}

	return meta
}

func (c *queryClient) GetAccount(accountID string, opts ...PayLoadMetaOption) *protocol.Query {
	meta := c.payloadMeta(opts...)

	return &protocol.Query{
		Payload: &protocol.Query_Payload{
			Query: &protocol.Query_Payload_GetAccount{
				GetAccount: &protocol.GetAccount{AccountId: accountID},
			},
			Meta: meta,
		},
	}
}

func (c *queryClient) GetSignatories(accountID string, opts ...PayLoadMetaOption) *protocol.Query {
	meta := c.payloadMeta(opts...)

	return &protocol.Query{
		Payload: &protocol.Query_Payload{
			Query: &protocol.Query_Payload_GetSignatories{
				GetSignatories: &protocol.GetSignatories{
					AccountId: accountID,
				},
			},
			Meta: meta,
		},
	}
}

func (c *queryClient) GetAccountTransactions(
	accountID string,
	txPaginationMeta *protocol.TxPaginationMeta,
	opts ...PayLoadMetaOption,
) *protocol.Query {
	meta := c.payloadMeta(opts...)

	return &protocol.Query{
		Payload: &protocol.Query_Payload{
			Query: &protocol.Query_Payload_GetAccountTransactions{
				GetAccountTransactions: &protocol.GetAccountTransactions{
					AccountId:      accountID,
					PaginationMeta: txPaginationMeta,
				},
			},
			Meta: meta,
		},
	}
}

func (c *queryClient) GetAccountAssetTransactions(
	accountID string,
	assetId string,
	txPaginationMeta *protocol.TxPaginationMeta,
	opts ...PayLoadMetaOption,
) *protocol.Query {
	meta := c.payloadMeta(opts...)

	return &protocol.Query{
		Payload: &protocol.Query_Payload{
			Query: &protocol.Query_Payload_GetAccountAssetTransactions{
				GetAccountAssetTransactions: &protocol.GetAccountAssetTransactions{
					AccountId:      accountID,
					AssetId:        assetId,
					PaginationMeta: txPaginationMeta,
				},
			},
			Meta: meta,
		},
	}
}

func (c *queryClient) GetTransactions(txHashes []string, opts ...PayLoadMetaOption) *protocol.Query {
	meta := c.payloadMeta(opts...)

	return &protocol.Query{
		Payload: &protocol.Query_Payload{
			Query: &protocol.Query_Payload_GetTransactions{
				GetTransactions: &protocol.GetTransactions{
					TxHashes: txHashes,
				},
			},
			Meta: meta,
		},
	}
}

func (c *queryClient) GetAccountAsset(
	accountID string,
	paginationMeta *protocol.AssetPaginationMeta,
	opts ...PayLoadMetaOption,
) *protocol.Query {
	meta := c.payloadMeta(opts...)

	return &protocol.Query{
		Payload: &protocol.Query_Payload{
			Query: &protocol.Query_Payload_GetAccountAssets{
				GetAccountAssets: &protocol.GetAccountAssets{
					AccountId:      accountID,
					PaginationMeta: paginationMeta,
				},
			},
			Meta: meta,
		},
	}
}

func (c *queryClient) GetAccountDetail(
	accountID string,
	key string,
	writer string,
	paginationMeta *protocol.AccountDetailPaginationMeta,
	opts ...PayLoadMetaOption,
) *protocol.Query {
	meta := c.payloadMeta(opts...)

	return &protocol.Query{
		Payload: &protocol.Query_Payload{
			Query: &protocol.Query_Payload_GetAccountDetail{
				GetAccountDetail: &protocol.GetAccountDetail{
					OptAccountId: &protocol.GetAccountDetail_AccountId{
						AccountId: accountID,
					},
					OptKey:         &protocol.GetAccountDetail_Key{Key: key},
					OptWriter:      &protocol.GetAccountDetail_Writer{Writer: writer},
					PaginationMeta: paginationMeta,
				},
			},
			Meta: meta,
		},
	}
}

func (c *queryClient) GetRoles(opts ...PayLoadMetaOption) *protocol.Query {
	meta := c.payloadMeta(opts...)

	return &protocol.Query{
		Payload: &protocol.Query_Payload{
			Query: &protocol.Query_Payload_GetRoles{
				GetRoles: &protocol.GetRoles{},
			},
			Meta: meta,
		},
	}
}

func (c *queryClient) GetRolePermissions(roleID string, opts ...PayLoadMetaOption) *protocol.Query {
	meta := c.payloadMeta(opts...)

	return &protocol.Query{
		Payload: &protocol.Query_Payload{
			Query: &protocol.Query_Payload_GetRolePermissions{
				GetRolePermissions: &protocol.GetRolePermissions{RoleId: roleID},
			},
			Meta: meta,
		},
	}
}

func (c *queryClient) GetAssetInfo(assetID string, opts ...PayLoadMetaOption) *protocol.Query {
	meta := c.payloadMeta(opts...)

	return &protocol.Query{
		Payload: &protocol.Query_Payload{
			Query: &protocol.Query_Payload_GetAssetInfo{
				GetAssetInfo: &protocol.GetAssetInfo{AssetId: assetID},
			},
			Meta: meta,
		},
	}
}

func (c *queryClient) GetPendingTransactions(
	txPaginationMeta *protocol.TxPaginationMeta,
	opts ...PayLoadMetaOption,
) *protocol.Query {
	meta := c.payloadMeta(opts...)

	return &protocol.Query{
		Payload: &protocol.Query_Payload{
			Query: &protocol.Query_Payload_GetPendingTransactions{
				GetPendingTransactions: &protocol.GetPendingTransactions{
					PaginationMeta: txPaginationMeta,
				},
			},
			Meta: meta,
		},
	}
}

func (c *queryClient) GetBlock(height uint64, opts ...PayLoadMetaOption) *protocol.Query {
	meta := c.payloadMeta(opts...)

	return &protocol.Query{
		Payload: &protocol.Query_Payload{
			Query: &protocol.Query_Payload_GetBlock{
				GetBlock: &protocol.GetBlock{Height: height},
			},
			Meta: meta,
		},
	}
}

func (c *queryClient) GetPeers(opts ...PayLoadMetaOption) *protocol.Query {
	meta := c.payloadMeta(opts...)

	return &protocol.Query{
		Payload: &protocol.Query_Payload{
			Query: &protocol.Query_Payload_GetPeers{
				GetPeers: &protocol.GetPeers{},
			},
			Meta: meta,
		},
	}
}

func (c *queryClient) GetEngineReceipts(txHash string, opts ...PayLoadMetaOption) *protocol.Query {
	meta := c.payloadMeta(opts...)

	return &protocol.Query{
		Payload: &protocol.Query_Payload{
			Query: &protocol.Query_Payload_GetEngineReceipts{
				GetEngineReceipts: &protocol.GetEngineReceipts{
					TxHash: txHash,
				},
			},
			Meta: meta,
		},
	}
}

func (c *queryClient) SendQuery(ctx context.Context, query *protocol.Query, privKeyHex string) (*protocol.QueryResponse, error) {
	sig, err := crypto.SignQuery(query, privKeyHex)
	if err != nil {
		return nil, err
	}

	query.Signature = sig

	reqCtx, cancel := context.WithTimeout(
		ctx,
		c.Timeout,
	)
	defer cancel()

	res, err := c.client.Find(reqCtx, query)
	if errRes, ok := res.Response.(*protocol.QueryResponse_ErrorResponse); ok {
		return nil, errors.New(errRes.ErrorResponse.String())
	} else if err != nil {
		return nil, err
	}

	return res, nil
}
