package query

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc"

	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
)

type QueryClient interface {
	SendQuery(ctx context.Context, query *pb.Query) (*pb.QueryResponse, error)
}

var _ QueryClient = (*queryClient)(nil)

type queryClient struct {
	Timeout time.Duration

	client   pb.QueryServiceV1Client
	callOpts []grpc.CallOption
	// TODO logger
}

func New(conn *grpc.ClientConn, timeout time.Duration, callOpts ...grpc.CallOption) QueryClient {
	return &queryClient{
		Timeout:  timeout,
		client:   pb.NewQueryServiceV1Client(conn),
		callOpts: callOpts,
	}
}

type PayLoadMetaOption func(*pb.QueryPayloadMeta)

func CreatorAccountId(id string) PayLoadMetaOption {
	return func(meta *pb.QueryPayloadMeta) {
		meta.CreatorAccountId = id
	}
}

func QueryCounter(counter uint64) PayLoadMetaOption {
	return func(meta *pb.QueryPayloadMeta) {
		meta.QueryCounter = counter
	}
}

func CreatedTime(t uint64) PayLoadMetaOption {
	return func(meta *pb.QueryPayloadMeta) {
		meta.CreatedTime = t
	}
}

func (c *queryClient) SendQuery(ctx context.Context, query *pb.Query) (*pb.QueryResponse, error) {
	reqCtx, cancel := context.WithTimeout(
		ctx,
		c.Timeout,
	)
	defer cancel()

	res, err := c.client.Find(reqCtx, query, c.callOpts...)
	if errRes, ok := res.Response.(*pb.QueryResponse_ErrorResponse); ok {
		return nil, errors.New(errRes.ErrorResponse.String())
	} else if err != nil {
		return nil, err
	}

	return res, nil
}

func GetAccount(accountID string, opts ...PayLoadMetaOption) *pb.Query {
	meta := payloadMeta(opts...)

	return &pb.Query{
		Payload: &pb.Query_Payload{
			Query: &pb.Query_Payload_GetAccount{
				GetAccount: &pb.GetAccount{AccountId: accountID},
			},
			Meta: meta,
		},
	}
}

func GetSignatories(accountID string, opts ...PayLoadMetaOption) *pb.Query {
	meta := payloadMeta(opts...)

	return &pb.Query{
		Payload: &pb.Query_Payload{
			Query: &pb.Query_Payload_GetSignatories{
				GetSignatories: &pb.GetSignatories{
					AccountId: accountID,
				},
			},
			Meta: meta,
		},
	}
}

func GetAccountTransactions(
	accountID string,
	txPaginationMeta *pb.TxPaginationMeta,
	opts ...PayLoadMetaOption,
) *pb.Query {
	meta := payloadMeta(opts...)

	return &pb.Query{
		Payload: &pb.Query_Payload{
			Query: &pb.Query_Payload_GetAccountTransactions{
				GetAccountTransactions: &pb.GetAccountTransactions{
					AccountId:      accountID,
					PaginationMeta: txPaginationMeta,
				},
			},
			Meta: meta,
		},
	}
}

func GetAccountAssetTransactions(
	accountID string,
	assetId string,
	txPaginationMeta *pb.TxPaginationMeta,
	opts ...PayLoadMetaOption,
) *pb.Query {
	meta := payloadMeta(opts...)

	return &pb.Query{
		Payload: &pb.Query_Payload{
			Query: &pb.Query_Payload_GetAccountAssetTransactions{
				GetAccountAssetTransactions: &pb.GetAccountAssetTransactions{
					AccountId:      accountID,
					AssetId:        assetId,
					PaginationMeta: txPaginationMeta,
				},
			},
			Meta: meta,
		},
	}
}

func GetTransactions(txHashes []string, opts ...PayLoadMetaOption) *pb.Query {
	meta := payloadMeta(opts...)

	return &pb.Query{
		Payload: &pb.Query_Payload{
			Query: &pb.Query_Payload_GetTransactions{
				GetTransactions: &pb.GetTransactions{
					TxHashes: txHashes,
				},
			},
			Meta: meta,
		},
	}
}

func GetAccountAsset(
	accountID string,
	paginationMeta *pb.AssetPaginationMeta,
	opts ...PayLoadMetaOption,
) *pb.Query {
	meta := payloadMeta(opts...)

	return &pb.Query{
		Payload: &pb.Query_Payload{
			Query: &pb.Query_Payload_GetAccountAssets{
				GetAccountAssets: &pb.GetAccountAssets{
					AccountId:      accountID,
					PaginationMeta: paginationMeta,
				},
			},
			Meta: meta,
		},
	}
}

func GetAccountDetail(
	accountID string,
	key string,
	writer string,
	paginationMeta *pb.AccountDetailPaginationMeta,
	opts ...PayLoadMetaOption,
) *pb.Query {
	meta := payloadMeta(opts...)

	return &pb.Query{
		Payload: &pb.Query_Payload{
			Query: &pb.Query_Payload_GetAccountDetail{
				GetAccountDetail: &pb.GetAccountDetail{
					OptAccountId: &pb.GetAccountDetail_AccountId{
						AccountId: accountID,
					},
					OptKey:         &pb.GetAccountDetail_Key{Key: key},
					OptWriter:      &pb.GetAccountDetail_Writer{Writer: writer},
					PaginationMeta: paginationMeta,
				},
			},
			Meta: meta,
		},
	}
}

func GetRoles(opts ...PayLoadMetaOption) *pb.Query {
	meta := payloadMeta(opts...)

	return &pb.Query{
		Payload: &pb.Query_Payload{
			Query: &pb.Query_Payload_GetRoles{
				GetRoles: &pb.GetRoles{},
			},
			Meta: meta,
		},
	}
}

func GetRolePermissions(roleID string, opts ...PayLoadMetaOption) *pb.Query {
	meta := payloadMeta(opts...)

	return &pb.Query{
		Payload: &pb.Query_Payload{
			Query: &pb.Query_Payload_GetRolePermissions{
				GetRolePermissions: &pb.GetRolePermissions{RoleId: roleID},
			},
			Meta: meta,
		},
	}
}

func GetAssetInfo(assetID string, opts ...PayLoadMetaOption) *pb.Query {
	meta := payloadMeta(opts...)

	return &pb.Query{
		Payload: &pb.Query_Payload{
			Query: &pb.Query_Payload_GetAssetInfo{
				GetAssetInfo: &pb.GetAssetInfo{AssetId: assetID},
			},
			Meta: meta,
		},
	}
}

func GetPendingTransactions(
	txPaginationMeta *pb.TxPaginationMeta,
	opts ...PayLoadMetaOption,
) *pb.Query {
	meta := payloadMeta(opts...)

	return &pb.Query{
		Payload: &pb.Query_Payload{
			Query: &pb.Query_Payload_GetPendingTransactions{
				GetPendingTransactions: &pb.GetPendingTransactions{
					PaginationMeta: txPaginationMeta,
				},
			},
			Meta: meta,
		},
	}
}

func GetBlock(height uint64, opts ...PayLoadMetaOption) *pb.Query {
	meta := payloadMeta(opts...)

	return &pb.Query{
		Payload: &pb.Query_Payload{
			Query: &pb.Query_Payload_GetBlock{
				GetBlock: &pb.GetBlock{Height: height},
			},
			Meta: meta,
		},
	}
}

func GetPeers(opts ...PayLoadMetaOption) *pb.Query {
	meta := payloadMeta(opts...)

	return &pb.Query{
		Payload: &pb.Query_Payload{
			Query: &pb.Query_Payload_GetPeers{
				GetPeers: &pb.GetPeers{},
			},
			Meta: meta,
		},
	}
}

func GetEngineReceipts(txHash string, opts ...PayLoadMetaOption) *pb.Query {
	meta := payloadMeta(opts...)

	return &pb.Query{
		Payload: &pb.Query_Payload{
			Query: &pb.Query_Payload_GetEngineReceipts{
				GetEngineReceipts: &pb.GetEngineReceipts{
					TxHash: txHash,
				},
			},
			Meta: meta,
		},
	}
}

func payloadMeta(opts ...PayLoadMetaOption) *pb.QueryPayloadMeta {
	meta := &pb.QueryPayloadMeta{
		QueryCounter: 1,
		CreatedTime:  uint64(time.Now().UnixNano() / int64(time.Millisecond)),
	}

	for _, opt := range opts {
		opt(meta)
	}

	return meta
}
