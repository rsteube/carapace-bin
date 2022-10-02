package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var toolchain_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed toolchains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolchain_listCmd).Standalone()

	toolchain_listCmd.Flags().BoolP("help", "h", false, "Prints help information")
	toolchain_listCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output with toolchain information")
	toolchainCmd.AddCommand(toolchain_listCmd)
}
