package command

import (
	"context"
	"encoding/hex"
	"errors"
	"io"
	"time"

	"google.golang.org/grpc"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/crypto"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
)

type CommandClient interface {
	AddAssetQuantity(assetID, amount string) *protocol.Command
	AddPeer(address, peerKey string, tlsCertificate *protocol.Peer_TlsCertificate) *protocol.Command
	AddSignatory(accountID, pubKey string) *protocol.Command
	AppendRole(accountID, roleName string) *protocol.Command
	CreateAccount(accountName, domainID, pubKey string) *protocol.Command
	CreateAsset(assetName, domainID string, precision uint32) *protocol.Command
	CreateDomain(defaultRole, domainID string) *protocol.Command
	CreateRole(roleName string, permissions []protocol.RolePermission) *protocol.Command
	DetachRole(accountId, roleName string) *protocol.Command
	GrantPermission(accountId string, permission protocol.GrantablePermission) *protocol.Command
	RemoveSignatory(accountId, pubKey string) *protocol.Command
	RevokePermission(accountId string, permission protocol.GrantablePermission) *protocol.Command
	SetAccountDetail(accountId, key, value string) *protocol.Command
	SetAccountQuorum(accountId string, quorum uint32) *protocol.Command
	SubtractAssetQuantity(assetId, amount string) *protocol.Command
	TransferAsset(srcAccountID, destAccountID, assetID, description, amount string) *protocol.Command
	RemovePeer(pubkey string) *protocol.Command
	CompareAndSetAccountDetail(accountID, key, value string, checkEmpty bool, optOldValue *protocol.CompareAndSetAccountDetail_OldValue) *protocol.Command
	SetSettingValue(key, value string) *protocol.Command
	CallEngine(caller string, callee *protocol.CallEngine_Callee, input string) *protocol.Command

	BuildPayload(cmds []*protocol.Command, opts ...PayLoadMetaOption) *protocol.Transaction_Payload_ReducedPayload
	BuildTransaction(reducedPayload *protocol.Transaction_Payload_ReducedPayload) *protocol.Transaction
	SendTransaction(ctx context.Context, tx *protocol.Transaction, privKeyHex string) (txHash string, err error)
	BuildBatchTransactions(txs []*protocol.Transaction, batchType protocol.Transaction_Payload_BatchMeta_BatchType) (*protocol.TxList, error)
	SendBatchTransaction(ctx context.Context, txList *protocol.TxList, privKeyHex string) ([]string, error)

	TxStatus(ctx context.Context, txHash string) (*protocol.ToriiResponse, error)
	TxStatusStream(ctx context.Context, txHash string) (*protocol.ToriiResponse, error)
}

var _ CommandClient = (*commandClient)(nil)

type commandClient struct {
	Timeout time.Duration

	client   protocol.CommandServiceV1Client
	callOpts []grpc.CallOption
	// TODO logger
}

func New(conn *grpc.ClientConn, timeout time.Duration, callOpts ...grpc.CallOption) CommandClient {
	return &commandClient{
		Timeout:  timeout,
		client:   protocol.NewCommandServiceV1Client(conn),
		callOpts: callOpts,
	}
}

type PayLoadMetaOption func(*protocol.Transaction_Payload_ReducedPayload)

func CreatorAccountId(id string) PayLoadMetaOption {
	return func(meta *protocol.Transaction_Payload_ReducedPayload) {
		meta.CreatorAccountId = id
	}
}

func Quorum(quorum uint32) PayLoadMetaOption {
	return func(meta *protocol.Transaction_Payload_ReducedPayload) {
		meta.Quorum = quorum
	}
}

func CreatedTime(t uint64) PayLoadMetaOption {
	return func(meta *protocol.Transaction_Payload_ReducedPayload) {
		meta.CreatedTime = t
	}
}

func (c *commandClient) AddAssetQuantity(assetID, amount string) *protocol.Command {
	return &protocol.Command{
		Command: &protocol.Command_AddAssetQuantity{
			AddAssetQuantity: &protocol.AddAssetQuantity{
				AssetId: assetID,
				Amount:  amount,
			},
		}}
}

func (c *commandClient) AddPeer(address, peerKey string, tlsCertificate *protocol.Peer_TlsCertificate) *protocol.Command {
	return &protocol.Command{
		Command: &protocol.Command_AddPeer{
			AddPeer: &protocol.AddPeer{
				Peer: &protocol.Peer{
					Address:     address,
					PeerKey:     peerKey,
					Certificate: tlsCertificate,
				},
			},
		}}
}

func (c *commandClient) AddSignatory(accountID, pubKey string) *protocol.Command {
	return &protocol.Command{
		Command: &protocol.Command_AddSignatory{
			AddSignatory: &protocol.AddSignatory{
				AccountId: accountID,
				PublicKey: pubKey,
			},
		}}
}

func (c *commandClient) AppendRole(accountID, roleName string) *protocol.Command {
	return &protocol.Command{
		Command: &protocol.Command_AppendRole{
			AppendRole: &protocol.AppendRole{
				AccountId: accountID,
				RoleName:  roleName,
			},
		}}
}

func (c *commandClient) CreateAccount(accountName, domainID, pubKey string) *protocol.Command {
	return &protocol.Command{
		Command: &protocol.Command_CreateAccount{
			CreateAccount: &protocol.CreateAccount{
				AccountName: accountName,
				DomainId:    domainID,
				PublicKey:   pubKey,
			},
		}}
}

func (c *commandClient) CreateAsset(assetName, domainID string, precision uint32) *protocol.Command {
	return &protocol.Command{
		Command: &protocol.Command_CreateAsset{
			CreateAsset: &protocol.CreateAsset{
				AssetName: assetName,
				DomainId:  domainID,
				Precision: precision,
			},
		}}
}

func (c *commandClient) CreateDomain(defaultRole, domainID string) *protocol.Command {
	return &protocol.Command{
		Command: &protocol.Command_CreateDomain{
			CreateDomain: &protocol.CreateDomain{
				DefaultRole: defaultRole,
				DomainId:    domainID,
			},
		}}
}

func (c *commandClient) CreateRole(roleName string, permissions []protocol.RolePermission) *protocol.Command {
	return &protocol.Command{
		Command: &protocol.Command_CreateRole{
			CreateRole: &protocol.CreateRole{
				RoleName:    roleName,
				Permissions: permissions,
			},
		}}
}

func (c *commandClient) DetachRole(accountId, roleName string) *protocol.Command {
	return &protocol.Command{
		Command: &protocol.Command_DetachRole{
			DetachRole: &protocol.DetachRole{
				AccountId: accountId,
				RoleName:  roleName,
			},
		}}
}

func (c *commandClient) GrantPermission(accountId string, permission protocol.GrantablePermission) *protocol.Command {
	return &protocol.Command{
		Command: &protocol.Command_GrantPermission{
			GrantPermission: &protocol.GrantPermission{
				AccountId:  accountId,
				Permission: permission,
			},
		}}
}

func (c *commandClient) RemoveSignatory(accountId, pubKey string) *protocol.Command {
	return &protocol.Command{
		Command: &protocol.Command_RemoveSignatory{
			RemoveSignatory: &protocol.RemoveSignatory{
				AccountId: accountId,
				PublicKey: pubKey,
			},
		}}
}

func (c *commandClient) RevokePermission(accountId string, permission protocol.GrantablePermission) *protocol.Command {
	return &protocol.Command{
		Command: &protocol.Command_RevokePermission{
			RevokePermission: &protocol.RevokePermission{
				AccountId:  accountId,
				Permission: permission,
			},
		}}
}

func (c *commandClient) SetAccountDetail(accountId, key, value string) *protocol.Command {
	return &protocol.Command{
		Command: &protocol.Command_SetAccountDetail{
			SetAccountDetail: &protocol.SetAccountDetail{
				AccountId: accountId,
				Key:       key,
				Value:     value,
			},
		}}
}

func (c *commandClient) SetAccountQuorum(accountId string, quorum uint32) *protocol.Command {
	return &protocol.Command{
		Command: &protocol.Command_SetAccountQuorum{
			SetAccountQuorum: &protocol.SetAccountQuorum{
				AccountId: accountId,
				Quorum:    quorum,
			},
		}}
}

func (c *commandClient) SubtractAssetQuantity(assetId, amount string) *protocol.Command {
	return &protocol.Command{
		Command: &protocol.Command_SubtractAssetQuantity{
			SubtractAssetQuantity: &protocol.SubtractAssetQuantity{
				AssetId: assetId,
				Amount:  amount,
			},
		}}
}

func (c *commandClient) TransferAsset(srcAccountID, destAccountID, assetID, description, amount string) *protocol.Command {
	return &protocol.Command{
		Command: &protocol.Command_TransferAsset{
			TransferAsset: &protocol.TransferAsset{
				SrcAccountId:  srcAccountID,
				DestAccountId: destAccountID,
				AssetId:       assetID,
				Description:   description,
				Amount:        amount,
			},
		}}
}

func (c *commandClient) RemovePeer(pubkey string) *protocol.Command {
	return &protocol.Command{
		Command: &protocol.Command_RemovePeer{
			RemovePeer: &protocol.RemovePeer{
				PublicKey: pubkey,
			},
		}}
}

func (c *commandClient) CompareAndSetAccountDetail(accountID, key, value string, checkEmpty bool, optOldValue *protocol.CompareAndSetAccountDetail_OldValue) *protocol.Command {
	return &protocol.Command{
		Command: &protocol.Command_CompareAndSetAccountDetail{
			CompareAndSetAccountDetail: &protocol.CompareAndSetAccountDetail{
				AccountId:   accountID,
				Key:         key,
				Value:       value,
				OptOldValue: optOldValue,
				CheckEmpty:  checkEmpty,
			},
		}}
}

func (c *commandClient) SetSettingValue(key, value string) *protocol.Command {
	return &protocol.Command{
		Command: &protocol.Command_SetSettingValue{
			SetSettingValue: &protocol.SetSettingValue{
				Key:   key,
				Value: value,
			},
		}}
}

func (c *commandClient) CallEngine(caller string, callee *protocol.CallEngine_Callee, input string) *protocol.Command {
	return &protocol.Command{
		Command: &protocol.Command_CallEngine{
			CallEngine: &protocol.CallEngine{
				Type:      protocol.CallEngine_kSolidity,
				Caller:    caller,
				OptCallee: callee,
				Input:     input,
			},
		}}
}

func (c *commandClient) BuildPayload(cmds []*protocol.Command, opts ...PayLoadMetaOption) *protocol.Transaction_Payload_ReducedPayload {
	payload := &protocol.Transaction_Payload_ReducedPayload{
		Commands:    cmds,
		CreatedTime: uint64(time.Now().UnixNano() / int64(time.Millisecond)),
		Quorum:      1,
	}

	for _, opt := range opts {
		opt(payload)
	}

	return payload
}

func (c *commandClient) BuildTransaction(reducedPayload *protocol.Transaction_Payload_ReducedPayload) *protocol.Transaction {
	tx := &protocol.Transaction{
		Payload: &protocol.Transaction_Payload{
			ReducedPayload: reducedPayload,
		},
	}

	return tx
}

func (c *commandClient) SendTransaction(ctx context.Context, tx *protocol.Transaction, privKeyHex string) (string, error) {
	sigs, err := crypto.SignTransaction(tx, privKeyHex)
	if err != nil {
		return "", err
	}

	tx.Signatures = sigs

	reqCtx, cancel := context.WithTimeout(
		ctx,
		c.Timeout,
	)
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

func (c *commandClient) BuildBatchTransactions(
	txs []*protocol.Transaction,
	batchType protocol.Transaction_Payload_BatchMeta_BatchType,
) (*protocol.TxList, error) {
	var reducedHashes []string
	for _, tx := range txs {
		hash, err := crypto.Hash(tx.GetPayload().ReducedPayload)
		if err != nil {
			return nil, err
		}
		reducedHashes = append(reducedHashes, hex.EncodeToString(hash))
	}

	for _, tx := range txs {
		tx.GetPayload().OptionalBatchMeta = &protocol.Transaction_Payload_Batch{
			Batch: &protocol.Transaction_Payload_BatchMeta{
				Type:          batchType,
				ReducedHashes: reducedHashes,
			},
		}
	}

	return &protocol.TxList{
		Transactions: txs,
	}, nil
}

func (c *commandClient) SendBatchTransaction(ctx context.Context, txList *protocol.TxList, privKeyHex string) ([]string, error) {
	for _, tx := range txList.Transactions {
		sigs, err := crypto.SignTransaction(tx, privKeyHex)
		if err != nil {
			return nil, err
		}

		tx.Signatures = sigs
	}

	reqCtx, cancel := context.WithTimeout(
		ctx,
		c.Timeout,
	)
	defer cancel()

	if _, err := c.client.ListTorii(reqCtx, txList, c.callOpts...); err != nil {
		return nil, err
	}

	var txHashList []string
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

func (c *commandClient) TxStatus(ctx context.Context, txHash string) (*protocol.ToriiResponse, error) {
	reqCtx, cancel := context.WithTimeout(
		ctx,
		c.Timeout,
	)
	defer cancel()

	res, err := c.client.Status(reqCtx, &protocol.TxStatusRequest{TxHash: txHash}, c.callOpts...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *commandClient) TxStatusStream(ctx context.Context, txHash string) (*protocol.ToriiResponse, error) {
	reqCtx, cancel := context.WithTimeout(
		ctx,
		c.Timeout,
	)
	defer cancel()

	stream, err := c.client.StatusStream(reqCtx, &protocol.TxStatusRequest{TxHash: txHash}, c.callOpts...)
	if err != nil {
		return nil, err
	}

	var res *protocol.ToriiResponse

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

	if res.TxStatus != protocol.TxStatus_COMMITTED {
		return nil, errors.New(res.String())
	}

	return res, nil
}
