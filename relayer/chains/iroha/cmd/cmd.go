package cmd

import (
	"github.com/hyperledger-labs/yui-relayer/config"
	"github.com/spf13/cobra"
)

func IrohaCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "iroha",
		Short: "manage iroha configurations",
	}

	cmd.AddCommand(
		configCmd(ctx),
		transactionCmd(ctx),
		apiCmd(ctx),
	)

	return cmd
}
