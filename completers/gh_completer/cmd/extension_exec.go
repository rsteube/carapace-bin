package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var extension_execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute an installed extension",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extension_execCmd).Standalone()
	extensionCmd.AddCommand(extension_execCmd)

	carapace.Gen(extensionCmd).PositionalCompletion(
		action.ActionExtensions(),
	)
}