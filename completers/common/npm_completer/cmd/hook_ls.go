package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var hook_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List registry hooks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_lsCmd).Standalone()

	hookCmd.AddCommand(hook_lsCmd)

	carapace.Gen(hook_lsCmd).PositionalCompletion(
		action.ActionPackageNames(hook_lsCmd),
	)
}
