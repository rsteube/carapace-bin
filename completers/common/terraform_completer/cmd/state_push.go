package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var state_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Update remote state from a local state file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_pushCmd).Standalone()

	stateCmd.AddCommand(state_pushCmd)

	carapace.Gen(state_pushCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
