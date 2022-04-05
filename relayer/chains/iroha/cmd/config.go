package cmd

import (
	"encoding/json"
	"os"
	"path"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/hyperledger-labs/yui-relayer/config"
	"github.com/hyperledger-labs/yui-relayer/utils"
	"github.com/spf13/cobra"

	"github.com/datachainlab/iroha-ibc-modules/relayer/chains/iroha"
)

func configCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "manage configuration file",
	}

	cmd.AddCommand(
		setContractAddressCmd(ctx),
	)

	return cmd
}

func setContractAddressCmd(ctx *config.Context) *cobra.Command {
	const (
		flagIbcHostAddress            = "ibc-host"
		flagIbcHandlerAddress         = "ibc-handler"
		flagIrohaIcs20BankAddress     = "iroha-ics20-bank"
		flagIrohaIcs20TransferAddress = "iroha-ics20-transfer"
	)
	cmd := &cobra.Command{
		Use:  "set-contract [chain-id]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			chainID := args[0]

			home, err := cmd.Flags().GetString(flags.FlagHome)
			if err != nil {
				return err
			}

			cfgPath := path.Join(home, "config", "config.yaml")
			if _, err := os.Stat(cfgPath); err != nil {
				return err
			}

			for i, proverConfig := range ctx.Config.Chains {
				var cfg iroha.ChainConfig

				if err := proverConfig.Init(ctx.Codec); err != nil {
					return err
				} else if chain, err := proverConfig.Build(); err != nil {
					return err
				} else if chain.ChainID() != chainID {
					continue
				}

				if err := json.Unmarshal(proverConfig.Chain, &cfg); err != nil {
					return err
				}

				if v, err := cmd.Flags().GetString(flagIbcHostAddress); err != nil {
					return err
				} else {
					cfg.IbcHostAddress = v
				}

				if v, err := cmd.Flags().GetString(flagIbcHandlerAddress); err != nil {
					return err
				} else {
					cfg.IbcHandlerAddress = v
				}

				if v, err := cmd.Flags().GetString(flagIrohaIcs20BankAddress); err != nil {
					return err
				} else {
					cfg.IrohaIcs20BankAddress = v
				}

				if v, err := cmd.Flags().GetString(flagIrohaIcs20TransferAddress); err != nil {
					return err
				} else {
					cfg.IrohaIcs20TransferAddress = v
				}

				if cbz, err := utils.MarshalJSONAny(ctx.Codec, &cfg); err != nil {
					return err
				} else {
					ctx.Config.Chains[i].Chain = cbz
				}

				bz, err := config.MarshalJSON(*ctx.Config)
				if err != nil {
					return err
				}

				f, err := os.Create(cfgPath)
				if err != nil {
					return err
				}
				defer f.Close()

				if _, err = f.Write(bz); err != nil {
					return err
				}
			}

			return nil
		},
	}

	cmd.Flags().String(flagIbcHostAddress, "", "the address of the IBCHost contract")
	cmd.Flags().String(flagIbcHandlerAddress, "", "the address of the IBCHandler contract")
	cmd.Flags().String(flagIrohaIcs20BankAddress, "", "the address of the IrohaICS20Bank contract")
	cmd.Flags().String(flagIrohaIcs20TransferAddress, "", "the address of the IrohaICS20TransferBank contract")

	return cmd
}
