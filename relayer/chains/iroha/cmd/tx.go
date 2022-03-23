package cmd

import (
	"context"
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
		Use:  "send-transfer [path-name] [chain-id] [src-account-id] [dest-account-id] [asset-id] [description] [amount]",
		Args: cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			pathName := args[0]
			chainID := args[1]
			srcAccountID := args[2]
			destAccountID := args[3]
			assetID := args[4]
			description := args[5]
			amount := args[6]

			// find source port id and source channel id
			var sourcePort, sourceChannel string
			if path, err := ctx.Config.Paths.Get(pathName); err != nil {
				return err
			} else if pathEnd := path.End(chainID); pathEnd.ChainID != chainID {
				return fmt.Errorf("PathEnd not found for pathName:%s and chainID:%s", pathName, chainID)
			} else {
				sourcePort = pathEnd.PortID
				sourceChannel = pathEnd.ChannelID
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

			// get latest block number and determine timeout height
			var timeoutHeight uint64
			if bn, err := getLatestBlockNumber(cmd.Context(), rpcCli); err != nil {
				return err
			} else {
				timeoutHeight = bn + 1000
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
		Use:  "burn [chain-id] [bank-account-id]",
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

			// submit burn tx
			tx, err := contract.Transact(transactOpts(cmd.Context(), bankAccountID), "burn")
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
		Use:  "mint [chain-id] [bank-account-id]",
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

			// submit mint tx
			tx, err := contract.Transact(transactOpts(cmd.Context(), bankAccountID), "mint")
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
