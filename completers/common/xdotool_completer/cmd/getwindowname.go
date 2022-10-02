package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/xdotool_completer/cmd/action"
	"github.com/spf13/cobra"
)

var getwindownameCmd = &cobra.Command{
	Use:   "getwindowname",
	Short: "Output the name of a given window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getwindownameCmd).Standalone()

	rootCmd.AddCommand(getwindownameCmd)

	carapace.Gen(getwindownameCmd).PositionalCompletion(
		action.ActionWindows(),
	)
}
