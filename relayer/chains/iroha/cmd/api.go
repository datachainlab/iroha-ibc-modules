package cmd

import (
	"fmt"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"
	"github.com/hyperledger-labs/yui-relayer/config"
	"github.com/spf13/cobra"
)

func apiCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api",
		Short: "call iroha api",
	}

	cmd.AddCommand(
		addAssetQuantityCmd(ctx),
		subtractAssetQuantityCmd(ctx),
		transferAssetCmd(ctx),
		getAccountAssetCmd(ctx),
	)

	return cmd
}

func addAssetQuantityCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "add-asset-quantity [chain-id] [signer-account-id] [signer-key-hex] [asset-id] [amount]",
		Args: cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			chainID := args[0]
			signerAccountID := args[1]
			signerKeyHex := args[2]
			assetID := args[3]
			amount := args[4]

			// find chain config
			cfg, err := findIrohaChainConfig(ctx, chainID)
			if err != nil {
				return err
			}

			// connect to irohad
			conn, err := dialToIrohad(cfg.ToriiAddr)
			if err != nil {
				return err
			}

			// build tx
			tx := command.BuildTransaction(
				command.BuildPayload(
					[]*protocol.Command{
						command.AddAssetQuantity(assetID, amount),
					},
					command.CreatorAccountId(signerAccountID),
				),
			)

			// send tx
			if _, err := sendIrohaTx(cmd.Context(), conn, tx, signerKeyHex); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}

func subtractAssetQuantityCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "subtract-asset-quantity [chain-id] [signer-account-id] [signer-key-hex] [asset-id] [amount]",
		Args: cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			chainID := args[0]
			signerAccountID := args[1]
			signerKeyHex := args[2]
			assetID := args[3]
			amount := args[4]

			// find chain config
			cfg, err := findIrohaChainConfig(ctx, chainID)
			if err != nil {
				return err
			}

			// connect to irohad
			conn, err := dialToIrohad(cfg.ToriiAddr)
			if err != nil {
				return err
			}

			// build tx
			tx := command.BuildTransaction(
				command.BuildPayload(
					[]*protocol.Command{
						command.SubtractAssetQuantity(assetID, amount),
					},
					command.CreatorAccountId(signerAccountID),
				),
			)

			// send tx
			if _, err := sendIrohaTx(cmd.Context(), conn, tx, signerKeyHex); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}

func transferAssetCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "transfer-asset [chain-id] [signer-account-id] [signer-key-hex] [src-account-id] [dest-account-id] [asset-id] [description] [amount]",
		Args: cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			chainID := args[0]
			signerAccountID := args[1]
			signerKeyHex := args[2]
			srcAccountID := args[3]
			destAccountID := args[4]
			assetID := args[5]
			description := args[6]
			amount := args[7]

			// find chain config
			cfg, err := findIrohaChainConfig(ctx, chainID)
			if err != nil {
				return err
			}

			// connect to irohad
			conn, err := dialToIrohad(cfg.ToriiAddr)
			if err != nil {
				return err
			}

			// build tx
			tx := command.BuildTransaction(
				command.BuildPayload(
					[]*protocol.Command{
						command.TransferAsset(srcAccountID, destAccountID, assetID, description, amount),
					},
					command.CreatorAccountId(signerAccountID),
				),
			)

			// send tx
			if _, err := sendIrohaTx(cmd.Context(), conn, tx, signerKeyHex); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}

func getAccountAssetCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get-account-asset [chain-id] [signer-key-hex] [account-id] [asset-id]",
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			chainID := args[0]
			signerKeyHex := args[1]
			accountID := args[2]
			assetID := args[3]

			// find chain config
			cfg, err := findIrohaChainConfig(ctx, chainID)
			if err != nil {
				return err
			}

			// connect to irohad
			conn, err := dialToIrohad(cfg.ToriiAddr)
			if err != nil {
				return err
			}

			// build query
			q := query.GetAccountAsset(accountID, nil, query.CreatorAccountId(accountID))

			// send query
			if res, err := sendIrohaQuery(cmd.Context(), conn, q, signerKeyHex); err != nil {
				return err
			} else {
				balance := "0"
				assets := res.GetAccountAssetsResponse().AccountAssets
				for _, a := range assets {
					if a.AssetId == assetID {
						balance = a.Balance
					}
				}
				fmt.Println(balance)
				return nil
			}
		},
	}

	return cmd
}
