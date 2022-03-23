package cmd

import (
	"context"
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
			cfg, err := findIrohaChainConfig(ctx, chainID)
			if err != nil {
				return err
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
