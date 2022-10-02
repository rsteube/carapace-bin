package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stateCmd = &cobra.Command{
	Use:   "state",
	Short: "Advanced state management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stateCmd).Standalone()

	rootCmd.AddCommand(stateCmd)
}
