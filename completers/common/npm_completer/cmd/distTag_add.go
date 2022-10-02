package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var distTag_addCmd = &cobra.Command{
	Use:   "add",
	Short: "tag specified version of the package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(distTag_addCmd).Standalone()

	distTagCmd.AddCommand(distTag_addCmd)

	carapace.Gen(distTag_addCmd).PositionalCompletion(
		action.ActionPackages(distTagCmd),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionPackageTags(distTagCmd, strings.Split(c.Args[0], "@")[0])
		}),
	)
}
