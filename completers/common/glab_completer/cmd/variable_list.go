package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var variable_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List project or group variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(variable_listCmd).Standalone()
	variable_listCmd.Flags().StringP("group", "g", "", "List group variables")

	variableCmd.AddCommand(variable_listCmd)

	carapace.Gen(variable_listCmd).FlagCompletion(carapace.ActionMap{
		"group": action.ActionGroups(variable_listCmd),
	})
}
