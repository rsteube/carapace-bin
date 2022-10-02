package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/circleci_completer/cmd/action"
	"github.com/spf13/cobra"
)

var context_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the named context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_deleteCmd).Standalone()
	context_deleteCmd.Flags().BoolP("force", "f", false, "Delete the context without asking for confirmation.")
	contextCmd.AddCommand(context_deleteCmd)

	// TODO org/context
	carapace.Gen(context_deleteCmd).PositionalCompletion(
		action.ActionVcsTypes(),
	)
}
