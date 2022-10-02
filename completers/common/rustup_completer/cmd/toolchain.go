package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var toolchainCmd = &cobra.Command{
	Use:   "toolchain",
	Short: "Modify or query the installed toolchains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolchainCmd).Standalone()

	toolchainCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.AddCommand(toolchainCmd)
}
