package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var setScriptCmd = &cobra.Command{
	Use:   "set-script",
	Short: "Set tasks in the scripts section of package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setScriptCmd).Standalone()
	addWorkspaceFlags(setScriptCmd)

	rootCmd.AddCommand(setScriptCmd)

	carapace.Gen(setScriptCmd).PositionalCompletion(
		action.ActionScripts(setScriptCmd),
	)
}
