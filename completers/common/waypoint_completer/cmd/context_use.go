package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/waypoint_completer/cmd/action"
	"github.com/spf13/cobra"
)

var context_useCmd = &cobra.Command{
	Use:   "use",
	Short: "Set the default context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_useCmd).Standalone()

	addGlobalOptions(context_useCmd)

	contextCmd.AddCommand(context_useCmd)

	carapace.Gen(context_useCmd).PositionalCompletion(
		action.ActionContexts(),
	)
}
