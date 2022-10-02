package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/waypoint_completer/cmd/action"
	"github.com/spf13/cobra"
)

var context_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Output context info",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_inspectCmd).Standalone()

	context_inspectCmd.Flags().Bool("json", false, "Output information in JSON format The default is false.")

	addGlobalOptions(context_inspectCmd)

	contextCmd.AddCommand(context_inspectCmd)

	carapace.Gen(context_inspectCmd).PositionalCompletion(
		action.ActionContexts(),
	)
}
