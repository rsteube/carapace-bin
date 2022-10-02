package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var sideloadCmd = &cobra.Command{
	Use:   "sideload",
	Short: "sideload the given full OTA package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sideloadCmd).Standalone()

	rootCmd.AddCommand(sideloadCmd)

	carapace.Gen(sideloadCmd).PositionalCompletion(
		carapace.ActionFiles(".zip"),
	)
}
