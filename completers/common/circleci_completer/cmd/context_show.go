package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/circleci_completer/cmd/action"
	"github.com/spf13/cobra"
)

var context_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_showCmd).Standalone()
	contextCmd.AddCommand(context_showCmd)

	// TODO org-name/context-name/secret-name
	carapace.Gen(context_showCmd).PositionalCompletion(
		action.ActionVcsTypes(),
	)
}
