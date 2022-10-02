package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var providersCmd = &cobra.Command{
	Use:   "providers",
	Short: "Show the providers required for this configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(providersCmd).Standalone()

	rootCmd.AddCommand(providersCmd)

	carapace.Gen(providersCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
