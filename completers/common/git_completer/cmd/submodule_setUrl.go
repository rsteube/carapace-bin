package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var submodule_setUrlCmd = &cobra.Command{
	Use:   "set-url",
	Short: "Sets the URL of the specified submodule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(submodule_setUrlCmd).Standalone()

	submoduleCmd.AddCommand(submodule_setUrlCmd)

	carapace.Gen(submodule_setUrlCmd).PositionalCompletion(
		git.ActionSubmoduleNames(),
	)
}
