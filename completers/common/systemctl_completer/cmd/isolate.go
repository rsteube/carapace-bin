package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var isolateCmd = &cobra.Command{
	Use:   "isolate",
	Short: "Start one unit and stop all others",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(isolateCmd).Standalone()

	rootCmd.AddCommand(isolateCmd)

	carapace.Gen(isolateCmd).PositionalCompletion(
		action.ActionUnits(),
	)
}
