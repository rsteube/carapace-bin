package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var toolchain_linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Create a custom toolchain by symlinking to a directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolchain_linkCmd).Standalone()

	toolchain_linkCmd.Flags().BoolP("help", "h", false, "Prints help information")
	toolchainCmd.AddCommand(toolchain_linkCmd)

	carapace.Gen(toolchain_linkCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionDirectories(),
	)
}
