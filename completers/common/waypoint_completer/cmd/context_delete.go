package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/waypoint_completer/cmd/action"
	"github.com/spf13/cobra"
)

var context_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_deleteCmd).Standalone()

	context_deleteCmd.Flags().Bool("all", false, "Delete all contexts.")

	addGlobalOptions(context_deleteCmd)

	contextCmd.AddCommand(context_deleteCmd)

	carapace.Gen(context_deleteCmd).PositionalCompletion(
		action.ActionContexts(),
	)
}
