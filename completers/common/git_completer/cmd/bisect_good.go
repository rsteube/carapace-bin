package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var bisect_goodCmd = &cobra.Command{
	Use:     "good",
	Aliases: []string{"old"},
	Short:   "mark know-good revisions",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bisect_goodCmd).Standalone()
	bisectCmd.AddCommand(bisect_goodCmd)

	carapace.Gen(bisect_goodCmd).PositionalAnyCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
