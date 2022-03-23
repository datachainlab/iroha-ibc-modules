package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
	"github.com/datachainlab/iroha-ibc-modules/relayer/chains/iroha"
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
			var cfg iroha.ChainConfig
			{
				var found bool
				for _, config := range ctx.Config.Chains {
					if err := config.Init(ctx.Codec); err != nil {
						return err
					} else if chain, err := config.Build(); err != nil {
						return err
					} else if chain.ChainID() == chainID {
						if err := json.Unmarshal(config.Chain, &cfg); err != nil {
							return err
						}
						found = true
						break
					}
				}
				if !found {
					return fmt.Errorf("chain config not found for chain_id:%s", chainID)
				}
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
