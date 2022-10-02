package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var submodule_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the registered submodule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(submodule_updateCmd).Standalone()
	// TODO find flag descriptions
	submodule_updateCmd.Flags().Bool("checkout", false, "Checkout commit of superproject as detached HEAD")
	submodule_updateCmd.Flags().String("depth", "", "Create a shallow clone")
	submodule_updateCmd.Flags().BoolP("force", "f", false, "throw away local changes")
	submodule_updateCmd.Flags().Bool("init", false, "Initialize all submodules")
	submodule_updateCmd.Flags().String("jobs", "", "Clone new submodules in parallel")
	submodule_updateCmd.Flags().Bool("merge", false, "Merge commit of superproject into submodule")
	submodule_updateCmd.Flags().BoolP("no-Fetch", "N", false, "Don't fetch from remote")
	submodule_updateCmd.Flags().Bool("no-recommend-shallow", false, "Ignore configured shallow recommendation")
	submodule_updateCmd.Flags().Bool("rebase", false, "Rebase current branch onto commit of superproject")
	submodule_updateCmd.Flags().Bool("recommend-shallow", false, "Initial clone will use recommended configuration")
	submodule_updateCmd.Flags().Bool("recursive", false, "Traverse submodules recursively")
	submodule_updateCmd.Flags().String("reference", "", "reference repository")
	submodule_updateCmd.Flags().Bool("remote", false, "Use submodules's remote tracking branch to update")

	submoduleCmd.AddCommand(submodule_updateCmd)

	carapace.Gen(submodule_updateCmd).FlagCompletion(carapace.ActionMap{
		"reference": carapace.ActionDirectories(),
	})

	carapace.Gen(submodule_updateCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionSubmodulePaths().Invoke(c).Filter(c.Args).ToA()
		}),
	)

}
