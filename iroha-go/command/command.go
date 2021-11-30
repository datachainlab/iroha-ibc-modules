package command

import (
	"context"
	"encoding/hex"
	"errors"
	"io"
	"time"

	"google.golang.org/grpc"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/crypto"
	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
)

type CommandClient interface {
	SendTransaction(ctx context.Context, tx *pb.Transaction) (txHash string, err error)
	SendBatchTransaction(ctx context.Context, txList *pb.TxList) ([]string, error)

	TxStatus(ctx context.Context, txHash string) (*pb.ToriiResponse, error)
	TxStatusStream(ctx context.Context, txHash string) (*pb.ToriiResponse, error)
}

var _ CommandClient = (*commandClient)(nil)

type commandClient struct {
	Timeout time.Duration

	client   pb.CommandServiceV1Client
	callOpts []grpc.CallOption
	// TODO logger
}

func New(conn *grpc.ClientConn, timeout time.Duration, callOpts ...grpc.CallOption) CommandClient {
	return &commandClient{
		Timeout:  timeout,
		client:   pb.NewCommandServiceV1Client(conn),
		callOpts: callOpts,
	}
}

type PayLoadMetaOption func(*pb.Transaction_Payload_ReducedPayload)

func CreatorAccountId(id string) PayLoadMetaOption {
	return func(meta *pb.Transaction_Payload_ReducedPayload) {
		meta.CreatorAccountId = id
	}
}

func Quorum(quorum uint32) PayLoadMetaOption {
	return func(meta *pb.Transaction_Payload_ReducedPayload) {
		meta.Quorum = quorum
	}
}

func CreatedTime(t uint64) PayLoadMetaOption {
	return func(meta *pb.Transaction_Payload_ReducedPayload) {
		meta.CreatedTime = t
	}
}

func (c *commandClient) SendTransaction(ctx context.Context, tx *pb.Transaction) (string, error) {
	reqCtx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()

	if _, err := c.client.Torii(reqCtx, tx, c.callOpts...); err != nil {
		return "", err
	}

	bz, err := crypto.Hash(tx.Payload)
	if err != nil {
		return "", err
	}

	txHash := hex.EncodeToString(bz)
	return txHash, nil
}

func (c *commandClient) SendBatchTransaction(ctx context.Context, txList *pb.TxList) ([]string, error) {
	reqCtx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()

	if _, err := c.client.ListTorii(reqCtx, txList, c.callOpts...); err != nil {
		return nil, err
	}

	txHashList := make([]string, 0, len(txList.Transactions))
	for _, tx := range txList.Transactions {
		bz, err := crypto.Hash(tx.Payload)
		if err != nil {
			return txHashList, err
		}

		txHash := hex.EncodeToString(bz)
		txHashList = append(txHashList, txHash)
	}

	return txHashList, nil
}

func (c *commandClient) TxStatus(ctx context.Context, txHash string) (*pb.ToriiResponse, error) {
	reqCtx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()

	res, err := c.client.Status(reqCtx, &pb.TxStatusRequest{TxHash: txHash}, c.callOpts...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *commandClient) TxStatusStream(ctx context.Context, txHash string) (*pb.ToriiResponse, error) {
	reqCtx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()

	stream, err := c.client.StatusStream(reqCtx, &pb.TxStatusRequest{TxHash: txHash}, c.callOpts...)
	if err != nil {
		return nil, err
	}

	var res *pb.ToriiResponse

	for {
		status, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		res = status
	}

	if res.TxStatus != pb.TxStatus_COMMITTED {
		return nil, errors.New(res.String())
	}

	return res, nil
}

func AddAssetQuantity(assetID, amount string) *pb.Command {
	return &pb.Command{
		Command: &pb.Command_AddAssetQuantity{
			AddAssetQuantity: &pb.AddAssetQuantity{
				AssetId: assetID,
				Amount:  amount,
			},
		}}
}

func AddPeer(address, peerKey string, tlsCertificate *pb.Peer_TlsCertificate) *pb.Command {
	return &pb.Command{
		Command: &pb.Command_AddPeer{
			AddPeer: &pb.AddPeer{
				Peer: &pb.Peer{
					Address:     address,
					PeerKey:     peerKey,
					Certificate: tlsCertificate,
				},
			},
		}}
}

func AddSignatory(accountID, pubKey string) *pb.Command {
	return &pb.Command{
		Command: &pb.Command_AddSignatory{
			AddSignatory: &pb.AddSignatory{
				AccountId: accountID,
				PublicKey: pubKey,
			},
		}}
}

func AppendRole(accountID, roleName string) *pb.Command {
	return &pb.Command{
		Command: &pb.Command_AppendRole{
			AppendRole: &pb.AppendRole{
				AccountId: accountID,
				RoleName:  roleName,
			},
		}}
}

func CreateAccount(accountName, domainID, pubKey string) *pb.Command {
	return &pb.Command{
		Command: &pb.Command_CreateAccount{
			CreateAccount: &pb.CreateAccount{
				AccountName: accountName,
				DomainId:    domainID,
				PublicKey:   pubKey,
			},
		}}
}

func CreateAsset(assetName, domainID string, precision uint32) *pb.Command {
	return &pb.Command{
		Command: &pb.Command_CreateAsset{
			CreateAsset: &pb.CreateAsset{
				AssetName: assetName,
				DomainId:  domainID,
				Precision: precision,
			},
		}}
}

func CreateDomain(defaultRole, domainID string) *pb.Command {
	return &pb.Command{
		Command: &pb.Command_CreateDomain{
			CreateDomain: &pb.CreateDomain{
				DefaultRole: defaultRole,
				DomainId:    domainID,
			},
		}}
}

func CreateRole(roleName string, permissions []pb.RolePermission) *pb.Command {
	return &pb.Command{
		Command: &pb.Command_CreateRole{
			CreateRole: &pb.CreateRole{
				RoleName:    roleName,
				Permissions: permissions,
			},
		}}
}

func DetachRole(accountId, roleName string) *pb.Command {
	return &pb.Command{
		Command: &pb.Command_DetachRole{
			DetachRole: &pb.DetachRole{
				AccountId: accountId,
				RoleName:  roleName,
			},
		}}
}

func GrantPermission(accountId string, permission pb.GrantablePermission) *pb.Command {
	return &pb.Command{
		Command: &pb.Command_GrantPermission{
			GrantPermission: &pb.GrantPermission{
				AccountId:  accountId,
				Permission: permission,
			},
		}}
}

func RemoveSignatory(accountId, pubKey string) *pb.Command {
	return &pb.Command{
		Command: &pb.Command_RemoveSignatory{
			RemoveSignatory: &pb.RemoveSignatory{
				AccountId: accountId,
				PublicKey: pubKey,
			},
		}}
}

func RevokePermission(accountId string, permission pb.GrantablePermission) *pb.Command {
	return &pb.Command{
		Command: &pb.Command_RevokePermission{
			RevokePermission: &pb.RevokePermission{
				AccountId:  accountId,
				Permission: permission,
			},
		}}
}

func SetAccountDetail(accountId, key, value string) *pb.Command {
	return &pb.Command{
		Command: &pb.Command_SetAccountDetail{
			SetAccountDetail: &pb.SetAccountDetail{
				AccountId: accountId,
				Key:       key,
				Value:     value,
			},
		}}
}

func SetAccountQuorum(accountId string, quorum uint32) *pb.Command {
	return &pb.Command{
		Command: &pb.Command_SetAccountQuorum{
			SetAccountQuorum: &pb.SetAccountQuorum{
				AccountId: accountId,
				Quorum:    quorum,
			},
		}}
}

func SubtractAssetQuantity(assetId, amount string) *pb.Command {
	return &pb.Command{
		Command: &pb.Command_SubtractAssetQuantity{
			SubtractAssetQuantity: &pb.SubtractAssetQuantity{
				AssetId: assetId,
				Amount:  amount,
			},
		}}
}

func TransferAsset(srcAccountID, destAccountID, assetID, description, amount string) *pb.Command {
	return &pb.Command{
		Command: &pb.Command_TransferAsset{
			TransferAsset: &pb.TransferAsset{
				SrcAccountId:  srcAccountID,
				DestAccountId: destAccountID,
				AssetId:       assetID,
				Description:   description,
				Amount:        amount,
			},
		}}
}

func RemovePeer(pubkey string) *pb.Command {
	return &pb.Command{
		Command: &pb.Command_RemovePeer{
			RemovePeer: &pb.RemovePeer{
				PublicKey: pubkey,
			},
		}}
}

func CompareAndSetAccountDetail(
	accountID, key, value string,
	checkEmpty bool,
	optOldValue *pb.CompareAndSetAccountDetail_OldValue,
) *pb.Command {
	return &pb.Command{
		Command: &pb.Command_CompareAndSetAccountDetail{
			CompareAndSetAccountDetail: &pb.CompareAndSetAccountDetail{
				AccountId:   accountID,
				Key:         key,
				Value:       value,
				OptOldValue: optOldValue,
				CheckEmpty:  checkEmpty,
			},
		}}
}

func SetSettingValue(key, value string) *pb.Command {
	return &pb.Command{
		Command: &pb.Command_SetSettingValue{
			SetSettingValue: &pb.SetSettingValue{
				Key:   key,
				Value: value,
			},
		}}
}

func CallEngine(caller string, callee string, input string) *pb.Command {
	return &pb.Command{
		Command: &pb.Command_CallEngine{
			CallEngine: &pb.CallEngine{
				Caller: caller,
				Callee: callee,
				Input:  input,
			},
		}}
}

func BuildPayload(cmds []*pb.Command, opts ...PayLoadMetaOption) *pb.Transaction_Payload_ReducedPayload {
	payload := &pb.Transaction_Payload_ReducedPayload{
		Commands:    cmds,
		CreatedTime: uint64(time.Now().UnixNano() / int64(time.Millisecond)),
		Quorum:      1,
	}

	for _, opt := range opts {
		opt(payload)
	}

	return payload
}

func BuildTransaction(reducedPayload *pb.Transaction_Payload_ReducedPayload) *pb.Transaction {
	tx := &pb.Transaction{
		Payload: &pb.Transaction_Payload{
			ReducedPayload: reducedPayload,
		},
	}

	return tx
}

func BuildBatchTransactions(
	txs []*pb.Transaction,
	batchType pb.Transaction_Payload_BatchMeta_BatchType,
) (*pb.TxList, error) {
	var reducedHashes []string
	for _, tx := range txs {
		hash, err := crypto.Hash(tx.GetPayload().ReducedPayload)
		if err != nil {
			return nil, err
		}
		reducedHashes = append(reducedHashes, hex.EncodeToString(hash))
	}

	for _, tx := range txs {
		tx.GetPayload().OptionalBatchMeta = &pb.Transaction_Payload_Batch{
			Batch: &pb.Transaction_Payload_BatchMeta{
				Type:          batchType,
				ReducedHashes: reducedHashes,
			},
		}
	}

	return &pb.TxList{
		Transactions: txs,
	}, nil
}
