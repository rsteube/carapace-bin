package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/waypoint_completer/cmd/action"
	"github.com/spf13/cobra"
)

var context_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify server access with a context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_verifyCmd).Standalone()

	addGlobalOptions(context_verifyCmd)

	contextCmd.AddCommand(context_verifyCmd)

	carapace.Gen(context_verifyCmd).PositionalCompletion(
		action.ActionContexts(),
	)
}
