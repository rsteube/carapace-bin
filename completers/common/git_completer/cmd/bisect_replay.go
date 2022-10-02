package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bisect_replayCmd = &cobra.Command{
	Use:   "replay",
	Short: "replay bisect log",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bisect_replayCmd).Standalone()

	bisectCmd.AddCommand(bisect_replayCmd)

	carapace.Gen(bisect_replayCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
