package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var hook_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove registry hook",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_rmCmd).Standalone()

	hookCmd.AddCommand(hook_rmCmd)

	// TODO positional completion
}
