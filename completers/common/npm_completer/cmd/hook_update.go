package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var hook_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update registry hook",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_updateCmd).Standalone()

	hookCmd.AddCommand(hook_updateCmd)

	// TODO positional completion
}
