package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var submodule_setBranchCmd = &cobra.Command{
	Use:   "set-branch",
	Short: "Sets the default remote tracking branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(submodule_setBranchCmd).Standalone()
	submodule_setBranchCmd.Flags().StringP("branch", "b", "", "branch to track")
	submodule_setBranchCmd.Flags().Bool("default", false, "removes the branch configuration key")

	submoduleCmd.AddCommand(submodule_setBranchCmd)

	carapace.Gen(submodule_setBranchCmd).FlagCompletion(carapace.ActionMap{
		"branch": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				// TODO this should be the default but is actually not working
				// because git wants the `--branch` flag defined before positional args
				return git.ActionSubmoduleBranches(c.Args[0])
			}
			// workaround for problem above: simply merge branch names from all remotes
			return git.ActionSubmoduleBranches()
		}),
	})

	carapace.Gen(submodule_setBranchCmd).PositionalCompletion(
		git.ActionSubmoduleNames(),
	)
}
