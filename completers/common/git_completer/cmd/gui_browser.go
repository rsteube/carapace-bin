package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var browserCmd = &cobra.Command{
	Use:   "browser",
	Short: "Start a tree browser showing all files in the specified commit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(browserCmd).Standalone()

	guiCmd.AddCommand(browserCmd)

	carapace.Gen(browserCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
