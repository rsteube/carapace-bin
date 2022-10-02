package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/circleci_completer/cmd/action"
	"github.com/spf13/cobra"
)

var context_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_createCmd).Standalone()
	contextCmd.AddCommand(context_createCmd)

	// TODO org/context
	carapace.Gen(context_createCmd).PositionalCompletion(
		action.ActionVcsTypes(),
	)
}
