package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var owner_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all the users who have access to modify a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(owner_lsCmd).Standalone()

	ownerCmd.AddCommand(owner_lsCmd)

	carapace.Gen(owner_lsCmd).PositionalCompletion(
		action.ActionPackageNames(owner_lsCmd),
	)
}
