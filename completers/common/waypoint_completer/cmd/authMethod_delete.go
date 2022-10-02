package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/waypoint_completer/cmd/action"
	"github.com/spf13/cobra"
)

var authMethod_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a previously configured auth method",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(authMethod_deleteCmd).Standalone()

	authMethodCmd.AddCommand(authMethod_deleteCmd)

	carapace.Gen(authMethod_deleteCmd).PositionalCompletion(
		action.ActionAuthMethods(),
	)
}
