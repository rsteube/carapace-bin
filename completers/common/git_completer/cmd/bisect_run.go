package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bisect_runCmd = &cobra.Command{
	Use:   "run",
	Short: "run script to automatically bisect",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bisect_runCmd).Standalone()

	bisectCmd.AddCommand(bisect_runCmd)

	carapace.Gen(bisect_runCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
