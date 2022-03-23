package cmd

import (
	"context"
	"strconv"
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
		sendTransferCmd(ctx),
		burnCmd(ctx),
		mintCmd(ctx),
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
			contract, err := iroha.NewIrohaIcs20Bank(cfg.IrohaICS20BankAddress(), rpcCli)
			if err != nil {
				return err
			}

			// submit setBank tx
			tx, err := contract.Transact(
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

func sendTransferCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "send-transfer [chain-id] [src-account-id] [dest-account-id] [asset-id] [description] [amount] [source-port] [source-channel] [timeout-height]",
		Args: cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) error {
			chainID := args[0]
			srcAccountID := args[1]
			destAccountID := args[2]
			assetID := args[3]
			description := args[4]
			amount := args[5]
			// TODO: the following three parameters should be determined by this command itself
			sourcePort := args[6]
			sourceChannel := args[7]
			timeoutHeight, err := strconv.ParseUint(args[8], 10, 64)
			if err != nil {
				return err
			}

			// find chain config
			cfg, err := findIrohaChainConfig(ctx, chainID)
			if err != nil {
				return err
			}

			// create IrohaICS20Transfer client
			rpcCli, err := iroha.NewRPCClient(cfg.RpcAddr)
			if err != nil {
				return err
			}
			contract, err := iroha.NewIrohaIcs20Transfer(cfg.IrohaICS20TransferAddress(), rpcCli)
			if err != nil {
				return err
			}

			// submit sendTransfer tx
			tx, err := contract.Transact(
				transactOpts(cmd.Context(), srcAccountID),
				"sendTransfer",
				srcAccountID,
				destAccountID,
				assetID,
				description,
				amount,
				sourcePort,
				sourceChannel,
				timeoutHeight,
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

func burnCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "burn [chain-id] [bank-account-id] [burn-request-id]",
		Args: cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			chainID := args[0]
			bankAccountID := args[1]
			burnRequestID := ""
			if len(args) > 2 {
				burnRequestID = args[2]
			}

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
			contract, err := iroha.NewIrohaIcs20Bank(cfg.IrohaICS20BankAddress(), rpcCli)
			if err != nil {
				return err
			}

			// submit burn tx
			tx, err := contract.Transact(
				transactOpts(cmd.Context(), bankAccountID),
				"burn",
				burnRequestID,
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

func mintCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "mint [chain-id] [bank-account-id] [mint-request-id]",
		Args: cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			chainID := args[0]
			bankAccountID := args[1]
			mintRequestID := ""
			if len(args) > 2 {
				mintRequestID = args[2]
			}

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
			contract, err := iroha.NewIrohaIcs20Bank(cfg.IrohaICS20BankAddress(), rpcCli)
			if err != nil {
				return err
			}

			// submit mint tx
			tx, err := contract.Transact(
				transactOpts(cmd.Context(), bankAccountID),
				"mint",
				mintRequestID,
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
