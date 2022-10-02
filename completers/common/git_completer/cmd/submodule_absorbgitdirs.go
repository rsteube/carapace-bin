package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var submodule_absorbgitdirsCmd = &cobra.Command{
	Use:   "absorbgitdirs",
	Short: "Move submodules git dir to superproject",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(submodule_absorbgitdirsCmd).Standalone()

	submoduleCmd.AddCommand(submodule_absorbgitdirsCmd)

	carapace.Gen(submodule_absorbgitdirsCmd).PositionalAnyCompletion(
		git.ActionSubmodulePaths(),
	)
}
