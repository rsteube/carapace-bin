package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var hook_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add registry hook",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_addCmd).Standalone()
	hook_addCmd.Flags().String("type", "", "type of hook")

	hookCmd.AddCommand(hook_addCmd)

	carapace.Gen(hook_addCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("packages", "owners", "scopes"),
	})

	carapace.Gen(hook_addCmd).PositionalCompletion(
		action.ActionPackageNames(hook_addCmd),
	)
}
