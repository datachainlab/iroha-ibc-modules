package evm

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/burrow/execution/native"
	"github.com/hyperledger/burrow/permission"

	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
	evmCtx "github.com/datachainlab/iroha-ibc-modules/web3-gateway/evm/context"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db/consts"
)

var (
	serviceContract = native.New().MustContract("ServiceContract",
		`* acmstate.ReaderWriter for bridging EVM state and Iroha state.
			* @dev This interface describes the functions exposed by the native service contracts layer in burrow.
			`,
		native.Function{
			Comment: `
				* @notice Gets asset balance of an Iroha account
				* @param Account Iroha account ID
				* @param Asset asset ID
				* @return Asset balance of the Account
				`,
			PermFlag: permission.Call,
			F:        getAssetBalance,
		},
		native.Function{
			Comment: `
				* @notice Transfers a certain amount of asset from some source account to destination account
				* @param Src source account address
				* @param Dst destination account address
				* @param Description description of the transfer
				* @param Asset asset ID
				* @param Amount amount to transfer
				* @return 'true' if successful, 'false' otherwise
				`,
			PermFlag: permission.Call,
			F:        transferAsset,
		},
		native.Function{
			Comment: `
				* @notice Creates a new iroha account
				* @param Name account name
				* @param Domain domain of account
				* @param Key key of account
				* @return 'true' if successful, 'false' otherwise
				`,
			PermFlag: permission.Call,
			F:        createAccount,
		},
		native.Function{
			Comment: `
				* @notice Adds asset to iroha account
				* @param Asset name of asset
				* @param Amount mount of asset to be added
				* @return 'true' if successful, 'false' otherwise
				`,
			PermFlag: permission.Call,
			F:        addAssetQuantity,
		},
		native.Function{
			Comment: `
				* @notice Subtracts asset from iroha account
				* @param Asset name of asset
				* @param Amount amount of asset to be subtracted
				* @return 'true' if successful, 'false' otherwise
				`,
			PermFlag: permission.Call,
			F:        subtractAssetQuantity,
		},
		native.Function{
			Comment: `
				* @notice Sets account detail
				* @param Account account id to be used
				* @param Key key for the added info
				* @param Value value of added info
				* @return 'true' if successful, 'false' otherwise
				`,
			PermFlag: permission.Call,
			F:        setAccountDetail,
		},
		native.Function{
			Comment: `
				* @notice Gets account detail
				* @param Account account id to be used
				* @return details of the account
				`,
			PermFlag: permission.Call,
			F:        getAccountDetail,
		},
		native.Function{
			Comment: `
				* @notice Sets account quorum
				* @param Account account id to be used
				* @param Quorum quorum value to be set
				* @return 'true' if successful, 'false' otherwise
				`,
			PermFlag: permission.Call,
			F:        setAccountQuorum,
		},
		native.Function{
			Comment: `
				* @notice Adds a signatory to the account
				* @param Account account id in which signatory to be added
				* @param Key publicy key to be added as signatory
				* @return 'true' if successful, 'false' otherwise
				`,
			PermFlag: permission.Call,
			F:        addSignatory,
		},
		native.Function{
			Comment: `
				* @notice Adds a signatory to the account
				* @param Account account id in which signatory to be added
				* @param Key publicy key to be added as signatory
				* @return 'true' if successful, 'false' otherwise
				`,
			PermFlag: permission.Call,
			F:        removeSignatory,
		},
		native.Function{
			Comment: `
				* @notice Creates a domain
				* @param Domain name of domain to be created
				* @param Role default role for user created in domain
				* @return 'true' if successful, 'false' otherwise
				`,
			PermFlag: permission.Call,
			F:        createDomain,
		},
		native.Function{
			Comment: `
				* @notice Gets state of the account
				* @param Account account id to be used
				* @return state of the account
				`,
			PermFlag: permission.Call,
			F:        getAccount,
		},
		native.Function{
			Comment: `
				* @notice Creates an asset
				* @param Name name of asset to be created
				* @param Domain domain of the created asset
				* @param Precision precision of created asset
				* @return 'true' if successful, 'false' otherwise
				`,
			PermFlag: permission.Call,
			F:        createAsset,
		},
		native.Function{
			Comment: `
				* @notice Get signatories of the account
				* @param Account account to be used
				* @return signatories of the account
				`,
			PermFlag: permission.Call,
			F:        getSignatories,
		},
		native.Function{
			Comment: `
				* @notice Get Asset's info
				* @param Asset asset id to be used
				* @return details of the asset
				`,
			PermFlag: permission.Call,
			F:        getAssetInfo,
		},
		native.Function{
			Comment: `
				* @notice Updates Account role
				* @param Account name of account to be updated
				* @param Role new role of the account
				* @return 'true' if successful, 'false' otherwise
				`,
			PermFlag: permission.Call,
			F:        appendRole,
		},
		native.Function{
			Comment: `
				* @notice Removes account role
				* @param Account name of account to be updated
				* @param Role role of the account to be removed
				* @return 'true' if successful, 'false' otherwise
				`,
			PermFlag: permission.Call,
			F:        detachRole,
		},
		native.Function{
			Comment: `
				* @notice Adds a new peer
				* @param Address address of the new peer 
				* @param PeerKey key of the new peer
				* @return 'true' if successful, 'false' otherwise
				`,
			PermFlag: permission.Call,
			F:        addPeer,
		},
		native.Function{
			Comment: `
				* @notice Removes a peer
				* @param PeerKey key of the peer to be removed
				* @return 'true' if successful, 'false' otherwise
				`,
			PermFlag: permission.Call,
			F:        removePeer,
		},
		native.Function{
			Comment: `
				* @notice Gets all peers
				* @return details of the peers
				`,
			PermFlag: permission.Call,
			F:        getPeers,
		},
		native.Function{
			Comment: `
				* @notice Gets block 
				* @param Height height of block to be used
				* @return the block at the given height 
				`,
			PermFlag: permission.Call,
			F:        getBlock,
		},
		native.Function{
			Comment: `
				* @notice Gets all roles
				* @return details of the roles
				`,
			PermFlag: permission.Call,
			F:        getRoles,
		},
		native.Function{
			Comment: `
				* @notice Gets permissions of the role
				* @param Role role id to be used
				* @return permissions of the given role
				`,
			PermFlag: permission.Call,
			F:        getRolePermissions,
		},
		native.Function{
			Comment: `
				* @notice Get transactions of the account
				* @param Account account to be used
				* @param TxPaginationMeta`,
			PermFlag: permission.Call,
			F:        getAccountTransactions,
		},
		native.Function{
			Comment: `
				* @notice Get pending transactions of the account
				* @param TxPaginationMeta`,
			PermFlag: permission.Call,
			F:        getPendingTransactions,
		},
		native.Function{
			Comment: `
				* @notice Get account asset transactions of the account
				* @param account Id 
				* @param asset Id
				* @param TxPaginationMeta`,
			PermFlag: permission.Call,
			F:        getAccountAssetTransactions,
		},
		native.Function{
			Comment: `
				* @notice Grant Permission
				* @param account  
				* @param permission`,
			PermFlag: permission.Call,
			F:        grantPermission,
		},
		native.Function{
			Comment: `
				* @notice Revoke Permission
				* @param account  
				* @param permission`,
			PermFlag: permission.Call,
			F:        revokePermission,
		},
		native.Function{
			Comment: `
				* @notice Compare And Set Account Detail
				* @param account  
				* @param key
				* @param value
				* @param old_value
				* @param check_empty`,
			PermFlag: permission.Call,
			F:        compareAndSetAccountDetail,
		},
		native.Function{
			Comment: `
				* @notice Get Transactions
				* @param tx hashes`,
			PermFlag: permission.Call,
			F:        getTransactions,
		},
		native.Function{
			Comment: `
				* @notice Create Role
				* @param role name
				* @param permissions`,
			PermFlag: permission.Call,
			F:        createRole,
		},
	)
)

type getAssetBalanceArgs struct {
	Account string
	Asset   string
}

type getAssetBalanceRets struct {
	Result string
}

func getAssetBalance(ctx native.Context, args getAssetBalanceArgs) (getAssetBalanceRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return getAssetBalanceRets{}, err
	}

	balances, err := callCtx.Execer.GetAccountAssets(args.Account)
	if err != nil {
		return getAssetBalanceRets{}, err
	}

	value := "0"
	for _, v := range balances {
		if v.AssetID == args.Asset {
			value = strconv.FormatInt(v.Amount, 10)
			break
		}
	}

	_ = ctx.Logger.Trace.Log("function", "getAssetBalance",
		"account", args.Account,
		"asset", args.Asset,
		"value", value)

	return getAssetBalanceRets{Result: value}, nil
}

type transferAssetArgs struct {
	Src    string
	Dst    string
	Asset  string
	Desc   string
	Amount string
}

type transferAssetRets struct {
	Result bool
}

func transferAsset(ctx native.Context, args transferAssetArgs) (transferAssetRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return transferAssetRets{Result: false}, err
	}

	err = callCtx.Execer.TransferAsset(args.Src, args.Dst, args.Asset, args.Desc, args.Amount)
	if err != nil {
		return transferAssetRets{Result: false}, err
	}

	_ = ctx.Logger.Trace.Log("function", "transferAsset",
		"src", args.Src,
		"dst", args.Dst,
		"assetID", args.Asset,
		"description", args.Desc,
		"amount", args.Amount)

	return transferAssetRets{Result: true}, nil
}

type createAccountArgs struct {
	Name   string
	Domain string
	Key    string
}

type createAccountRets struct {
	Result bool
}

func createAccount(ctx native.Context, args createAccountArgs) (createAccountRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return createAccountRets{Result: false}, err
	}
	err = callCtx.Execer.CreateAccount(args.Name, args.Domain, args.Key)
	if err != nil {
		return createAccountRets{Result: false}, err
	}

	_ = ctx.Logger.Trace.Log("function", "createAccount",
		"name", args.Name,
		"domain", args.Domain,
		"key", args.Key)

	return createAccountRets{Result: true}, nil
}

type addAssetQuantityArgs struct {
	Asset  string
	Amount string
}

type addAssetQuantityRets struct {
	Result bool
}

func addAssetQuantity(ctx native.Context, args addAssetQuantityArgs) (addAssetQuantityRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return addAssetQuantityRets{Result: false}, err
	}
	err = callCtx.Execer.AddAssetQuantity(args.Asset, args.Amount)
	if err != nil {
		return addAssetQuantityRets{Result: false}, err
	}

	_ = ctx.Logger.Trace.Log("function", "addAssetQuantity",
		"asset", args.Asset,
		"amount", args.Amount)

	return addAssetQuantityRets{Result: true}, nil
}

type subtractAssetQuantityArgs struct {
	Asset  string
	Amount string
}

type subtractAssetQuantityRets struct {
	Result bool
}

func subtractAssetQuantity(ctx native.Context, args subtractAssetQuantityArgs) (subtractAssetQuantityRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return subtractAssetQuantityRets{Result: false}, err
	}
	err = callCtx.Execer.SubtractAssetQuantity(args.Asset, args.Amount)
	if err != nil {
		return subtractAssetQuantityRets{Result: false}, err
	}

	_ = ctx.Logger.Trace.Log("function", "subtractAssetQuantity",
		"asset", args.Asset,
		"amount", args.Amount)

	return subtractAssetQuantityRets{Result: true}, nil
}

type setAccountDetailArgs struct {
	Account string
	Key     string
	Value   string
}

type setAccountDetailRets struct {
	Result bool
}

func setAccountDetail(ctx native.Context, args setAccountDetailArgs) (setAccountDetailRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return setAccountDetailRets{Result: false}, err
	}
	err = callCtx.Execer.SetAccountDetail(args.Account, args.Key, args.Value)
	if err != nil {
		return setAccountDetailRets{Result: false}, err
	}

	_ = ctx.Logger.Trace.Log("function", "setAccountDetail",
		"account", args.Account,
		"key", args.Key,
		"value", args.Value)

	return setAccountDetailRets{Result: true}, nil
}

type getAccountDetailArgs struct {
}

type getAccountDetailRets struct {
	Result string
}

func getAccountDetail(ctx native.Context, args getAccountDetailArgs) (getAccountDetailRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return getAccountDetailRets{}, err
	}
	details, err := callCtx.Execer.GetAccountDetail()
	if err != nil {
		return getAccountDetailRets{}, err
	}

	_ = ctx.Logger.Trace.Log("function", "getAccountDetail")

	return getAccountDetailRets{Result: details}, nil
}

type setAccountQuorumArgs struct {
	Account string
	Quorum  string
}

type setAccountQuorumRets struct {
	Result bool
}

func setAccountQuorum(ctx native.Context, args setAccountQuorumArgs) (setAccountQuorumRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return setAccountQuorumRets{Result: false}, err
	}
	err = callCtx.Execer.SetAccountQuorum(args.Account, args.Quorum)
	if err != nil {
		return setAccountQuorumRets{Result: false}, err
	}

	_ = ctx.Logger.Trace.Log("function", "setAccountQuorum",
		"account", args.Account,
		"quorum", args.Quorum)

	return setAccountQuorumRets{Result: true}, nil
}

type addSignatoryArgs struct {
	Account string
	Key     string
}

type addSignatoryRets struct {
	Result bool
}

func addSignatory(ctx native.Context, args addSignatoryArgs) (addSignatoryRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return addSignatoryRets{Result: false}, err
	}
	err = callCtx.Execer.AddSignatory(args.Account, args.Key)
	if err != nil {
		return addSignatoryRets{Result: false}, err
	}

	_ = ctx.Logger.Trace.Log("function", "addSignatory",
		"account id", args.Account,
		"public key", args.Key)

	return addSignatoryRets{Result: true}, nil
}

type removeSignatoryArgs struct {
	Account string
	Key     string
}

type removeSignatoryRets struct {
	Result bool
}

func removeSignatory(ctx native.Context, args removeSignatoryArgs) (removeSignatoryRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return removeSignatoryRets{Result: false}, err
	}
	err = callCtx.Execer.RemoveSignatory(args.Account, args.Key)
	if err != nil {
		return removeSignatoryRets{Result: false}, err
	}

	_ = ctx.Logger.Trace.Log("function", "removeSignatory",
		"account id", args.Account,
		"public key", args.Key)

	return removeSignatoryRets{Result: true}, nil
}

type createDomainArgs struct {
	Domain string
	Role   string
}

type createDomainRets struct {
	Result bool
}

func createDomain(ctx native.Context, args createDomainArgs) (createDomainRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return createDomainRets{Result: false}, err
	}
	err = callCtx.Execer.CreateDomain(args.Domain, args.Role)
	if err != nil {
		return createDomainRets{Result: false}, err
	}

	_ = ctx.Logger.Trace.Log("function", "createDomain",
		"domain name", args.Domain,
		"default role", args.Role)

	return createDomainRets{Result: true}, nil
}

type getAccountArgs struct {
	Account string
}

type getAccountRets struct {
	Result string
}

func getAccount(ctx native.Context, args getAccountArgs) (getAccountRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return getAccountRets{}, err
	}

	ret, err := callCtx.Execer.GetAccount(args.Account)
	if err != nil {
		return getAccountRets{}, err
	}

	account := &pb.Account{
		AccountId: ret.AccountID,
		DomainId:  ret.DomainID,
		Quorum:    uint32(ret.Quorum),
		JsonData:  ret.Data,
	}

	_ = ctx.Logger.Trace.Log("function", "getAccount",
		"account", args.Account,
		"domain", account.GetDomainId(),
		"quorum", fmt.Sprint(account.GetQuorum()))
	result, err := json.Marshal(account)
	return getAccountRets{Result: string(result)}, nil
}

type createAssetArgs struct {
	Name      string
	Domain    string
	Precision string
}

type createAssetRets struct {
	Result bool
}

func createAsset(ctx native.Context, args createAssetArgs) (createAssetRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return createAssetRets{Result: false}, err
	}
	err = callCtx.Execer.CreateAsset(args.Name, args.Domain, args.Precision)
	if err != nil {
		return createAssetRets{Result: false}, err
	}

	_ = ctx.Logger.Trace.Log("function", "createAsset",
		"asset name", args.Name,
		"domain id", args.Domain,
		"precision", args.Precision)

	return createAssetRets{Result: true}, nil
}

type getSignatoriesArgs struct {
	Account string
}

type getSignatoriesRets struct {
	Keys []string
}

func getSignatories(ctx native.Context, args getSignatoriesArgs) (getSignatoriesRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return getSignatoriesRets{}, err
	}
	signatory, err := callCtx.Execer.GetSignatories(args.Account)
	if err != nil {
		return getSignatoriesRets{}, err
	}

	_ = ctx.Logger.Trace.Log("function", "getSignatories",
		"account", args.Account,
		"key", signatory)

	return getSignatoriesRets{Keys: signatory}, nil
}

type getAssetInfoArgs struct {
	Asset string
}

type getAssetInfoRets struct {
	Result string
}

func getAssetInfo(ctx native.Context, args getAssetInfoArgs) (getAssetInfoRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return getAssetInfoRets{}, err
	}
	ret, err := callCtx.Execer.GetAssetInfo(args.Asset)
	if err != nil {
		return getAssetInfoRets{}, err
	}

	asset := &pb.Asset{
		AssetId:   args.Asset,
		DomainId:  ret.DomainID,
		Precision: uint32(ret.Precision),
	}

	_ = ctx.Logger.Trace.Log("function", "getAssetInfo",
		"asset", args.Asset,
		"domain", asset.GetDomainId(),
		"precision", fmt.Sprint(asset.GetPrecision()))
	result, err := json.Marshal(asset)
	return getAssetInfoRets{Result: string(result)}, nil
}

type appendRoleArgs struct {
	Account string
	Role    string
}

type appendRoleRets struct {
	Result bool
}

func appendRole(ctx native.Context, args appendRoleArgs) (appendRoleRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return appendRoleRets{Result: false}, err
	}
	err = callCtx.Execer.AppendRole(args.Account, args.Role)
	if err != nil {
		return appendRoleRets{Result: false}, err
	}

	_ = ctx.Logger.Trace.Log("function", "appendRole",
		"account name", args.Account,
		"new role", args.Role)

	return appendRoleRets{Result: true}, nil
}

type detachRoleArgs struct {
	Account string
	Role    string
}

type detachRoleRets struct {
	Result bool
}

func detachRole(ctx native.Context, args detachRoleArgs) (detachRoleRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return detachRoleRets{Result: false}, err
	}
	err = callCtx.Execer.DetachRole(args.Account, args.Role)
	if err != nil {
		return detachRoleRets{Result: false}, err
	}

	_ = ctx.Logger.Trace.Log("function", "detachRole",
		"account name", args.Account,
		"removed role", args.Role)

	return detachRoleRets{Result: true}, nil
}

type addPeerArgs struct {
	Address string
	PeerKey string
}

type addPeerRets struct {
	Result bool
}

func addPeer(ctx native.Context, args addPeerArgs) (addPeerRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return addPeerRets{Result: false}, err
	}
	err = callCtx.Execer.AddPeer(args.Address, args.PeerKey)
	if err != nil {
		return addPeerRets{Result: false}, err
	}

	_ = ctx.Logger.Trace.Log("function", "addPeer",
		"peer address", args.Address,
		"peer key", args.PeerKey)

	return addPeerRets{Result: true}, nil
}

type removePeerArgs struct {
	PeerKey string
}

type removePeerRets struct {
	Result bool
}

func removePeer(ctx native.Context, args removePeerArgs) (removePeerRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return removePeerRets{Result: false}, err
	}
	err = callCtx.Execer.RemovePeer(args.PeerKey)
	if err != nil {
		return removePeerRets{Result: false}, err
	}

	_ = ctx.Logger.Trace.Log("function", "removePeer",
		"peer key", args.PeerKey)

	return removePeerRets{Result: true}, nil
}

type GrantPermissionArgs struct {
	AccountId  string
	Permission string
}

type GrantPermissionRets struct {
	Result bool
}

func grantPermission(ctx native.Context, args GrantPermissionArgs) (GrantPermissionRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return GrantPermissionRets{Result: false}, err
	}
	err = callCtx.Execer.GrantPermission(args.AccountId, args.Permission)
	if err != nil {
		return GrantPermissionRets{Result: false}, err
	}

	_ = ctx.Logger.Trace.Log("function", "GrantPermission",
		"account", args.AccountId, "Permission", args.Permission)

	return GrantPermissionRets{Result: true}, nil
}

type RevokePermissionArgs = GrantPermissionArgs
type RevokePermissionRets = GrantPermissionRets

func revokePermission(ctx native.Context, args RevokePermissionArgs) (RevokePermissionRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return RevokePermissionRets{Result: false}, err
	}
	err = callCtx.Execer.RevokePermission(args.AccountId, args.Permission)
	if err != nil {
		return RevokePermissionRets{Result: false}, err
	}

	_ = ctx.Logger.Trace.Log("function", "RevokePermission",
		"account", args.AccountId, "Permission", args.Permission)

	return RevokePermissionRets{Result: true}, nil
}

type compareAndSetAccountDetailArgs struct {
	AccountId  string
	Key        string
	Value      string
	OldValue   string
	CheckEmpty string
}

type compareAndSetAccountDetailRets struct {
	Result bool
}

func compareAndSetAccountDetail(ctx native.Context, args compareAndSetAccountDetailArgs) (compareAndSetAccountDetailRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return compareAndSetAccountDetailRets{Result: false}, err
	}
	err = callCtx.Execer.CompareAndSetAccountDetail(args.AccountId, args.Key, args.Value, args.OldValue, args.CheckEmpty)
	if err != nil {
		return compareAndSetAccountDetailRets{Result: false}, err
	}

	_ = ctx.Logger.Trace.Log("function", "CompareAndSetAccountDetail",
		"account", args.AccountId, "key", args.Key, "value", args.Value,
		"old value", args.OldValue, "check empty", args.CheckEmpty)

	return compareAndSetAccountDetailRets{Result: true}, nil
}

type createRoleArgs struct {
	RoleName    string
	Permissions string
}

type createRoleRets struct {
	Result bool
}

func createRole(ctx native.Context, args createRoleArgs) (createRoleRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return createRoleRets{Result: false}, err
	}
	err = callCtx.Execer.CreateRole(args.RoleName, args.Permissions)
	if err != nil {
		return createRoleRets{Result: false}, err
	}

	_ = ctx.Logger.Trace.Log("function", "CreateRole",
		"Role Name", args.RoleName, "Permissions", args.Permissions)

	return createRoleRets{Result: true}, nil
}

type getPeersArgs struct {
}

type getPeersRets struct {
	Result string
}

func getPeers(ctx native.Context, _ getPeersArgs) (getPeersRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return getPeersRets{}, err
	}
	ret, err := callCtx.Execer.GetPeers()
	if err != nil {
		return getPeersRets{}, err
	}

	peers := make([]*pb.Peer, 0, len(ret))
	for _, p := range ret {
		peer := &pb.Peer{
			PeerKey: p.PublicKey,
			Address: p.Address,
		}
		if p.TlsCertificate.Valid {
			peer.Certificate = &pb.Peer_TlsCertificate{
				TlsCertificate: p.TlsCertificate.String,
			}
		}
		peers = append(peers, peer)
	}

	_ = ctx.Logger.Trace.Log("function", "getPeers")
	result, err := json.Marshal(peers)
	return getPeersRets{Result: string(result)}, nil
}

type getBlockArgs struct {
	Height string
}

type getBlockRets struct {
	Result string
}

func getBlock(ctx native.Context, args getBlockArgs) (getBlockRets, error) {
	return getBlockRets{}, errors.New("not implemented: getBlock")
}

type getRolesArgs struct {
}

type getRolesRets struct {
	Result []string
}

func getRoles(ctx native.Context, _ getRolesArgs) (getRolesRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return getRolesRets{}, err
	}
	roles, err := callCtx.Execer.GetRoles()
	if err != nil {
		return getRolesRets{}, err
	}
	_ = ctx.Logger.Trace.Log("function", "getRoles")
	return getRolesRets{Result: roles}, nil
}

type getRolePermissionsArgs struct {
	Role string
}

type getRolePermissionsRets struct {
	Result string
}

func getRolePermissions(ctx native.Context, args getRolePermissionsArgs) (getRolePermissionsRets, error) {
	callCtx, err := evmCtx.LoadCallContext(ctx.CallParams)
	if err != nil {
		return getRolePermissionsRets{}, err
	}
	ret, err := callCtx.Execer.GetRolePermissions(args.Role)
	if err != nil {
		return getRolePermissionsRets{}, err
	}

	var permissions []pb.RolePermission
	lastIdx := consts.RolePermissionEnumLength - 1
	for i := lastIdx; i >= 0; i-- {
		if ret[i] == '1' {
			permissions = append(permissions, pb.RolePermission(lastIdx-i))
		}
	}

	_ = ctx.Logger.Trace.Log("function", "getRolePermissions",
		"role id", args.Role)

	result, err := json.Marshal(permissions)
	return getRolePermissionsRets{Result: string(result)}, nil
}

type GetAccountTransactionsArgs struct {
	Account       string
	PageSize      string
	FirstTxHash   string
	FirstTxTime   string
	LastTxTime    string
	FirstTxHeight string
	LastTxHeight  string
	Ordering      string
}

type getAccountTransactionsRets struct {
	Result string
}

func getAccountTransactions(ctx native.Context, args GetAccountTransactionsArgs) (getAccountTransactionsRets, error) {
	return getAccountTransactionsRets{}, errors.New("not implemented: getAccountTransactions")
}

type GetPendingTransactionsArgs struct {
	PageSize    string
	FirstTxHash string
	FirstTxTime string
	LastTxTime  string
	Ordering    string
}

type getPendingTransactionsRets struct {
	Result string
}

func getPendingTransactions(ctx native.Context, args GetPendingTransactionsArgs) (getPendingTransactionsRets, error) {
	return getPendingTransactionsRets{}, errors.New("not implemented: getPendingTransactions")
}

type GetAccountAssetTransactionsArgs struct {
	AccountId     string
	AssetId       string
	PageSize      string
	FirstTxHash   string
	FirstTxTime   string
	LastTxTime    string
	FirstTxHeight string
	LastTxHeight  string
	Ordering      string
}

type getAccountAssetTransactionsRets struct {
	Result string
}

func getAccountAssetTransactions(ctx native.Context, args GetAccountAssetTransactionsArgs) (getAccountAssetTransactionsRets, error) {
	return getAccountAssetTransactionsRets{}, errors.New("not implemented: getAccountAssetTransactions")
}

type GetTransactionsArgs struct {
	Hashes string
}

type getTransactionsRets struct {
	Result string
}

func getTransactions(ctx native.Context, args GetTransactionsArgs) (getTransactionsRets, error) {
	return getTransactionsRets{}, errors.New("not implemented: getTransactions")
}

func isNative(acc string) bool {
	return strings.ToLower(acc) == "a6abc17819738299b3b2c1ce46d55c74f04e290c"
}
