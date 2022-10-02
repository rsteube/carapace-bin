package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var submodule_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the status of the submodule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(submodule_statusCmd).Standalone()
	submodule_statusCmd.Flags().Bool("cached", false, "print the SHA-1 recorded in the superproject for each submodule")
	submodule_statusCmd.Flags().Bool("recursive", false, "recurse into nested submodules")

	submoduleCmd.AddCommand(submodule_statusCmd)

	carapace.Gen(submodule_statusCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionSubmodulePaths().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
