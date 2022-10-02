package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var bisect_resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "finish bisection search and go back to commit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bisect_resetCmd).Standalone()

	bisectCmd.AddCommand(bisect_resetCmd)

	carapace.Gen(bisect_resetCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
