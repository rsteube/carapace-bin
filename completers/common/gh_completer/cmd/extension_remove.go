package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var extension_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an installed extension",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extension_removeCmd).Standalone()
	extensionCmd.AddCommand(extension_removeCmd)

	carapace.Gen(extension_removeCmd).PositionalCompletion(
		action.ActionExtensions(),
	)
}
