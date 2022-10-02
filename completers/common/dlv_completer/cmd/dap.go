package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var dapCmd = &cobra.Command{
	Use:   "dap",
	Short: "Starts a headless TCP server communicating via Debug Adaptor Protocol (DAP).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dapCmd).Standalone()
	dapCmd.Flags().String("client-addr", "", "host:port where the DAP client is waiting for the DAP server to dial in")
	rootCmd.AddCommand(dapCmd)
}
