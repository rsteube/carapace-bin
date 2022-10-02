package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var state_pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull current state and output to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_pullCmd).Standalone()

	stateCmd.AddCommand(state_pullCmd)
}
