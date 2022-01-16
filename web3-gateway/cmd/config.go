package cmd

import (
	"bytes"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "display the configuration of the web3-gateway server",
	Run: func(cmd *cobra.Command, args []string) {
		var b bytes.Buffer

		enc := yaml.NewEncoder(&b)
		enc.SetIndent(2)

		cobra.CheckErr(enc.Encode(cfg))

		_, err := cmd.OutOrStdout().Write(b.Bytes())
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
