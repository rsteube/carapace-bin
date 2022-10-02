package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var distTag_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "clear a tag that is no longer in use from the package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(distTag_rmCmd).Standalone()

	distTagCmd.AddCommand(distTag_rmCmd)

	carapace.Gen(distTag_rmCmd).PositionalCompletion(
		action.ActionPackageNames(distTag_rmCmd),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionPackageTags(distTag_rmCmd, c.Args[0])
		}),
	)
}
