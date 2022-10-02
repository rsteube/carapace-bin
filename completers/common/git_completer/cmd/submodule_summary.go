package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var submodule_summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show commit summary between the given commit and working tree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(submodule_summaryCmd).Standalone()
	submodule_summaryCmd.Flags().Bool("cached", false, "")
	submodule_summaryCmd.Flags().Bool("files", false, "")
	submodule_summaryCmd.Flags().String("summary-limit", "", "")

	submoduleCmd.AddCommand(submodule_summaryCmd)

	carapace.Gen(submodule_summaryCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()), // TODO commmit of submodule??
	)

	carapace.Gen(submodule_summaryCmd).PositionalAnyCompletion(
		carapace.ActionFiles(), // TODO or submodule path??
	)
}
