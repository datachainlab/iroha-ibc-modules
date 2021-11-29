package evm

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/hyperledger/burrow/execution/native"
	"github.com/hyperledger/burrow/permission"

	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/acm"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db"
)

var (
	once            sync.Once
	callCtx         *callContext
	serviceContract = native.New().MustContract("ServiceContract",
		`* tcmstate.ReaderWriter for bridging EVM state and Iroha state.
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
		//native.Function{
		//	Comment: `
		//		* @notice Transfers a certain amount of asset from some source account to destination account
		//		* @param Src source account address
		//		* @param Dst destination account address
		//		* @param Description description of the transfer
		//		* @param Asset asset ID
		//		* @param Amount amount to transfer
		//		* @return 'true' if successful, 'false' otherwise
		//		`,
		//	PermFlag: permission.Call,
		//	F:        transferAsset,
		//},
		//native.Function{
		//	Comment: `
		//		* @notice Creates a new iroha account
		//		* @param Name account name
		//		* @param Domain domain of account
		//		* @param Key key of account
		//		* @return 'true' if successful, 'false' otherwise
		//		`,
		//	PermFlag: permission.Call,
		//	F:        createAccount,
		//},
		//native.Function{
		//	Comment: `
		//		* @notice Adds asset to iroha account
		//		* @param Asset name of asset
		//		* @param Amount mount of asset to be added
		//		* @return 'true' if successful, 'false' otherwise
		//		`,
		//	PermFlag: permission.Call,
		//	F:        addAssetQuantity,
		//},
		//native.Function{
		//	Comment: `
		//		* @notice Subtracts asset from iroha account
		//		* @param Asset name of asset
		//		* @param Amount amount of asset to be subtracted
		//		* @return 'true' if successful, 'false' otherwise
		//		`,
		//	PermFlag: permission.Call,
		//	F:        subtractAssetQuantity,
		//},
		//native.Function{
		//	Comment: `
		//		* @notice Sets account detail
		//		* @param Account account id to be used
		//		* @param Key key for the added info
		//		* @param Value value of added info
		//		* @return 'true' if successful, 'false' otherwise
		//		`,
		//	PermFlag: permission.Call,
		//	F:        setAccountDetail,
		//},
		native.Function{
			Comment: `
				* @notice Gets account detail
				* @param Account account id to be used
				* @return details of the account
				`,
			PermFlag: permission.Call,
			F:        getAccountDetail,
		},
		//native.Function{
		//	Comment: `
		//		* @notice Sets account quorum
		//		* @param Account account id to be used
		//		* @param Quorum quorum value to be set
		//		* @return 'true' if successful, 'false' otherwise
		//		`,
		//	PermFlag: permission.Call,
		//	F:        setAccountQuorum,
		//},
		//native.Function{
		//	Comment: `
		//		* @notice Adds a signatory to the account
		//		* @param Account account id in which signatory to be added
		//		* @param Key publicy key to be added as signatory
		//		* @return 'true' if successful, 'false' otherwise
		//		`,
		//	PermFlag: permission.Call,
		//	F:        addSignatory,
		//},
		//native.Function{
		//	Comment: `
		//		* @notice Adds a signatory to the account
		//		* @param Account account id in which signatory to be added
		//		* @param Key publicy key to be added as signatory
		//		* @return 'true' if successful, 'false' otherwise
		//		`,
		//	PermFlag: permission.Call,
		//	F:        removeSignatory,
		//},
		//native.Function{
		//	Comment: `
		//		* @notice Creates a domain
		//		* @param Domain name of domain to be created
		//		* @param Role default role for user created in domain
		//		* @return 'true' if successful, 'false' otherwise
		//		`,
		//	PermFlag: permission.Call,
		//	F:        createDomain,
		//},
		native.Function{
			Comment: `
				* @notice Gets state of the account
				* @param Account account id to be used
				* @return state of the account
				`,
			PermFlag: permission.Call,
			F:        getAccount,
		},
		//native.Function{
		//	Comment: `
		//		* @notice Creates an asset
		//		* @param Name name of asset to be created
		//		* @param Domain domain of the created asset
		//		* @param Precision precision of created asset
		//		* @return 'true' if successful, 'false' otherwise
		//		`,
		//	PermFlag: permission.Call,
		//	F:        createAsset,
		//},
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
		//native.Function{
		//	Comment: `
		//		* @notice Updates Account role
		//		* @param Account name of account to be updated
		//		* @param Role new role of the account
		//		* @return 'true' if successful, 'false' otherwise
		//		`,
		//	PermFlag: permission.Call,
		//	F:        appendRole,
		//},
		//native.Function{
		//	Comment: `
		//		* @notice Removes account role
		//		* @param Account name of account to be updated
		//		* @param Role role of the account to be removed
		//		* @return 'true' if successful, 'false' otherwise
		//		`,
		//	PermFlag: permission.Call,
		//	F:        detachRole,
		//},
		//native.Function{
		//	Comment: `
		//		* @notice Adds a new peer
		//		* @param Address address of the new peer
		//		* @param PeerKey key of the new peer
		//		* @return 'true' if successful, 'false' otherwise
		//		`,
		//	PermFlag: permission.Call,
		//	F:        addPeer,
		//},
		//native.Function{
		//	Comment: `
		//		* @notice Removes a peer
		//		* @param PeerKey key of the peer to be removed
		//		* @return 'true' if successful, 'false' otherwise
		//		`,
		//	PermFlag: permission.Call,
		//	F:        removePeer,
		//},
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
				* @param txPaginationMeta`,
			PermFlag: permission.Call,
			F:        getAccountTransactions,
		},
		native.Function{
			Comment: `
				* @notice Get pending transactions of the account
				* @param txPaginationMeta`,
			PermFlag: permission.Call,
			F:        getPendingTransactions,
		},
		native.Function{
			Comment: `
				* @notice Get account asset transactions of the account
				* @param account Id 
				* @param asset Id
				* @param txPaginationMeta`,
			PermFlag: permission.Call,
			F:        getAccountAssetTransactions,
		},
		//native.Function{
		//	Comment: `
		//		* @notice Grant Permission
		//		* @param account
		//		* @param permission`,
		//	PermFlag: permission.Call,
		//	F:       grantPermission,
		//},
		//native.Function{
		//	Comment: `
		//		* @notice Revoke Permission
		//		* @param account
		//		* @param permission`,
		//	PermFlag: permission.Call,
		//	F:       revokePermission,
		//},
		//native.Function{
		//	Comment: `
		//		* @notice Compare And Set Account Detail
		//		* @param account
		//		* @param key
		//		* @param value
		//		* @param old_value
		//		* @param check_empty`,
		//	PermFlag: permission.Call,
		//	F:       compareAndSetAccountDetail,
		//},
		native.Function{
			Comment: `
				* @notice Get Transactions
				* @param tx hashes`,
			PermFlag: permission.Call,
			F:        getTransactions,
		},
		//native.Function{
		//	Comment: `
		//		* @notice Create Role
		//		* @param role name
		//		* @param permissions`,
		//	PermFlag: permission.Call,
		//	F:        createRole,
		//},
	)
)

type callContext struct {
	queryClient      query.QueryClient
	dbClient         db.DBClient
	querierAccountID string
	keyStore         acm.KeyStore
}

func RegisterCallContext(
	queryClient query.QueryClient,
	dbClient db.DBClient,
	querierAccountID string,
	keyStore acm.KeyStore,
) *callContext {
	once.Do(func() {
		callCtx = &callContext{
			queryClient:      queryClient,
			dbClient:         dbClient,
			querierAccountID: querierAccountID,
			keyStore:         keyStore,
		}
	})

	return callCtx
}

func CallContext() *callContext {
	return callCtx
}

type getAssetBalanceArgs struct {
	Account string
	Asset   string
}

type getAssetBalanceRets struct {
	Result string
}

func getAssetBalance(ctx native.Context, args getAssetBalanceArgs) (getAssetBalanceRets, error) {
	res, err := callCtx.sendQuery(
		query.GetAccountAsset(
			args.Account,
			nil,
			query.CreatorAccountId(callCtx.querierAccountID),
		),
	)
	if err != nil {
		return getAssetBalanceRets{}, err
	}

	balances := res.GetAccountAssetsResponse().GetAccountAssets()

	value := "0"
	for _, v := range balances {
		if v.GetAssetId() == args.Asset {
			value = v.GetBalance()
			break
		}
	}

	ctx.Logger.Trace.Log("function", "getAssetBalance",
		"account", args.Account,
		"asset", args.Asset,
		"value", value)

	return getAssetBalanceRets{Result: value}, nil
}

type getAccountDetailArgs struct {
}

type getAccountDetailRets struct {
	Result string
}

func getAccountDetail(ctx native.Context, args getAccountDetailArgs) (getAccountDetailRets, error) {
	res, err := callCtx.sendQuery(
		query.GetAccountDetail(
			nil, nil, nil, nil,
			query.CreatorAccountId(callCtx.querierAccountID),
		))
	if err != nil {
		return getAccountDetailRets{}, err
	}

	details := res.GetAccountDetailResponse().GetDetail()

	ctx.Logger.Trace.Log("function", "getAccountDetail")

	return getAccountDetailRets{Result: details}, nil
}

type getAccountArgs struct {
	Account string
}

type getAccountRets struct {
	Result string
}

func getAccount(ctx native.Context, args getAccountArgs) (getAccountRets, error) {
	res, err := callCtx.sendQuery(
		query.GetAccount(
			args.Account,
			query.CreatorAccountId(callCtx.querierAccountID),
		),
	)
	if err != nil {
		return getAccountRets{}, err
	}

	account := res.GetAccountResponse().GetAccount()
	ctx.Logger.Trace.Log("function", "getAccount",
		"account", args.Account,
		"domain", account.GetDomainId(),
		"quorum", fmt.Sprint(account.GetQuorum()))
	result, err := json.Marshal(account)
	return getAccountRets{Result: string(result)}, nil
}

type getSignatoriesArgs struct {
	Account string
}

type getSignatoriesRets struct {
	Keys []string
}

func getSignatories(ctx native.Context, args getSignatoriesArgs) (getSignatoriesRets, error) {
	res, err := callCtx.sendQuery(
		query.GetSignatories(
			args.Account,
			query.CreatorAccountId(callCtx.querierAccountID),
		),
	)
	if err != nil {
		return getSignatoriesRets{}, err
	}

	signatories := res.GetSignatoriesResponse().GetKeys()

	ctx.Logger.Trace.Log("function", "getSignatories",
		"account", args.Account,
		"key", signatories)

	return getSignatoriesRets{Keys: signatories}, nil
}

type getAssetInfoArgs struct {
	Asset string
}

type getAssetInfoRets struct {
	Result string
}

func getAssetInfo(ctx native.Context, args getAssetInfoArgs) (getAssetInfoRets, error) {
	res, err := callCtx.sendQuery(
		query.GetAssetInfo(
			args.Asset,
			query.CreatorAccountId(callCtx.querierAccountID),
		),
	)
	if err != nil {
		return getAssetInfoRets{}, err
	}

	asset := res.GetAssetResponse().GetAsset()

	ctx.Logger.Trace.Log("function", "getAssetInfo",
		"asset", args.Asset,
		"domain", asset.GetDomainId(),
		"precision", fmt.Sprint(asset.GetPrecision()))
	result, err := json.Marshal(asset)
	return getAssetInfoRets{Result: string(result)}, nil
}

type getPeersArgs struct {
}

type getPeersRets struct {
	Result string
}

func getPeers(ctx native.Context, args getPeersArgs) (getPeersRets, error) {
	res, err := callCtx.sendQuery(
		query.GetPeers(
			query.CreatorAccountId(callCtx.querierAccountID),
		),
	)
	if err != nil {
		return getPeersRets{}, err
	}

	peers := res.GetPeersResponse().GetPeers()

	ctx.Logger.Trace.Log("function", "getPeers")
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
	height, err := strconv.ParseUint(args.Height, 10, 64)
	if err != nil {
		return getBlockRets{}, err
	}

	res, err := callCtx.sendQuery(
		query.GetBlock(
			height,
			query.CreatorAccountId(callCtx.querierAccountID),
		),
	)
	if err != nil {
		return getBlockRets{}, err
	}

	block := res.GetBlockResponse().GetBlock()

	ctx.Logger.Trace.Log("function", "getBlock",
		"block height", args.Height)
	result, err := json.Marshal(block)
	return getBlockRets{Result: string(result)}, nil
}

type getRolesArgs struct {
}

type getRolesRets struct {
	Result []string
}

func getRoles(ctx native.Context, args getRolesArgs) (getRolesRets, error) {
	res, err := callCtx.sendQuery(
		query.GetRoles(
			query.CreatorAccountId(callCtx.querierAccountID),
		))
	if err != nil {
		return getRolesRets{}, err
	}

	roles := res.GetRolesResponse().GetRoles()

	ctx.Logger.Trace.Log("function", "getRoles")

	return getRolesRets{Result: roles}, nil
}

type getRolePermissionsArgs struct {
	Role string
}

type getRolePermissionsRets struct {
	Result string
}

func getRolePermissions(ctx native.Context, args getRolePermissionsArgs) (getRolePermissionsRets, error) {
	res, err := callCtx.sendQuery(
		query.GetRolePermissions(
			args.Role,
			query.CreatorAccountId(callCtx.querierAccountID),
		))
	if err != nil {
		return getRolePermissionsRets{}, err
	}

	permissions := res.GetRolePermissionsResponse().GetPermissions()

	ctx.Logger.Trace.Log("function", "getRolePermissions",
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
	paginationMetaArg, err := makeTxPaginationMeta(
		&txPaginationMeta{PageSize: &args.PageSize, FirstTxHash: &args.PageSize, Ordering: &args.Ordering,
			FirstTxTime: &args.FirstTxTime, LastTxTime: &args.LastTxTime, FirstTxHeight: &args.FirstTxHeight, LastTxHeight: &args.LastTxHeight},
	)
	if err != nil {
		return getAccountTransactionsRets{}, err
	}
	res, err := callCtx.sendQuery(
		query.GetAccountTransactions(
			args.Account,
			&paginationMetaArg,
			query.CreatorAccountId(callCtx.querierAccountID),
		))
	if err != nil {
		return getAccountTransactionsRets{}, err
	}

	transactions := res.GetTransactionsResponse().GetTransactions()

	ctx.Logger.Trace.Log("function", "GetAccountTransactions",
		"account", args.Account)
	result, err := json.Marshal(transactions)
	return getAccountTransactionsRets{Result: string(result)}, nil
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
	paginationMetaArg, err := makeTxPaginationMeta(
		&txPaginationMeta{PageSize: &args.PageSize, FirstTxHash: &args.PageSize, Ordering: &args.Ordering,
			FirstTxTime: &args.FirstTxTime, LastTxTime: &args.LastTxTime},
	)
	res, err := callCtx.sendQuery(
		query.GetPendingTransactions(
			&paginationMetaArg,
			query.CreatorAccountId(callCtx.querierAccountID),
		))
	if err != nil {
		return getPendingTransactionsRets{}, err
	}

	transactions := res.GetPendingTransactionsPageResponse().GetTransactions()
	ctx.Logger.Trace.Log("function", "GetPendingTransactions")
	result, err := json.Marshal(transactions)
	return getPendingTransactionsRets{Result: string(result)}, nil
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
	paginationMetaArg, err := makeTxPaginationMeta(
		&txPaginationMeta{PageSize: &args.PageSize, FirstTxHash: &args.PageSize, Ordering: &args.Ordering,
			FirstTxTime: &args.FirstTxTime, LastTxTime: &args.LastTxTime, FirstTxHeight: &args.FirstTxHeight, LastTxHeight: &args.LastTxHeight},
	)
	res, err := callCtx.sendQuery(
		query.GetAccountAssetTransactions(
			args.AccountId,
			args.AssetId,
			&paginationMetaArg,
			query.CreatorAccountId(callCtx.querierAccountID),
		))
	if err != nil {
		return getAccountAssetTransactionsRets{}, err
	}

	transactions := res.GetTransactionsPageResponse().GetTransactions()
	ctx.Logger.Trace.Log("function", "GetAccountAssetTransactions", "account", args.AccountId, "asset", args.AssetId)
	result, err := json.Marshal(transactions)
	return getAccountAssetTransactionsRets{Result: string(result)}, nil
}

type GetTransactionsArgs struct {
	Hashes string
}

type getTransactionsRets struct {
	Result string
}

func getTransactions(ctx native.Context, args GetTransactionsArgs) (getTransactionsRets, error) {
	res, err := callCtx.sendQuery(
		query.GetTransactions(
			[]string{args.Hashes},
			query.CreatorAccountId(callCtx.querierAccountID),
		))
	if err != nil {
		return getTransactionsRets{}, err
	}

	transactions := res.GetTransactionsResponse().GetTransactions()
	ctx.Logger.Trace.Log("function", "GetTransactions", "hashes", args.Hashes)
	result, err := json.Marshal(transactions)
	return getTransactionsRets{Result: string(result)}, nil
}

func (c *callContext) sendQuery(q *pb.Query) (*pb.QueryResponse, error) {
	_, err := c.keyStore.SignQuery(q, c.querierAccountID)
	if err != nil {
		return nil, err
	}

	return c.queryClient.SendQuery(context.Background(), q)
}

func isNative(acc string) bool {
	return strings.ToLower(acc) == "a6abc17819738299b3b2c1ce46d55c74f04e290c"
}
