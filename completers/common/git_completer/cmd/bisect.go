package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var bisectCmd = &cobra.Command{
	Use:   "bisect",
	Short: "Use binary search to find the commit that introduced a bug",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bisectCmd).Standalone()

	rootCmd.AddCommand(bisectCmd)

	carapace.Gen(bisect_skipCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("...", func(c carapace.Context) carapace.Action {
			if len(c.Parts) > 2 {
				return carapace.ActionValues()
			}
			return git.ActionRefs(git.RefOption{}.Default())
		}),
	)
}
