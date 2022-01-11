package cmd

import (
	"github.com/spf13/cobra"

	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/rpc"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start the web3-gateway server",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(rpc.Serve(cfg))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
