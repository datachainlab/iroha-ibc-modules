package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/datachainlab/iroha-ibc-modules/relayer/chains/iroha"
	"github.com/hyperledger-labs/yui-relayer/config"
	"github.com/spf13/cobra"
)

func transactionCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tx",
		Short: "send eth tx to iroha (gateway)",
	}

	cmd.AddCommand(
		setBankCmd(ctx),
	)

	return cmd
}

func setBankCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "set-bank [chain-id] [bank-account-id]",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			chainID := args[0]
			bankAccountID := args[1]

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

			// create IrohaICS20Bank client
			rpcCli, err := iroha.NewRPCClient(cfg.RpcAddr)
			if err != nil {
				return err
			}
			bank, err := iroha.NewIrohaIcs20Bank(cfg.IrohaICS20BankAddress(), rpcCli)
			if err != nil {
				return err
			}

			// submit tx of setBank
			tx, err := bank.Transact(
				transactOpts(cmd.Context(), cfg.AccountId),
				"setBank",
				bankAccountID,
			)
			if err != nil {
				return err
			}

			// wait for receipt
			toCtx, cancel := context.WithTimeout(cmd.Context(), time.Minute)
			_, err = waitForReceipt(toCtx, rpcCli, tx.ID)
			defer cancel()

			return err
		},
	}

	return cmd
}
