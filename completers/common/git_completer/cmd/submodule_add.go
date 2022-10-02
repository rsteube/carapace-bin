package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var submodule_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add the given repository as a submodule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(submodule_addCmd).Standalone()
	submodule_addCmd.Flags().StringP("branch", "b", "", "branch of repository to add as submodule")
	submodule_addCmd.Flags().Bool("depth", false, "create a shallow clone of that depth")
	submodule_addCmd.Flags().Bool("dissociate", false, "use --reference only when cloning")
	submodule_addCmd.Flags().BoolP("force", "f", false, "alllow adding an otherwise ignored submodule path")
	submodule_addCmd.Flags().String("name", "", "set submodule name")
	submodule_addCmd.Flags().Bool("progress", false, "report progress status")
	submodule_addCmd.Flags().String("reference", "", "reference repository")

	submoduleCmd.AddCommand(submodule_addCmd)

	carapace.Gen(submodule_addCmd).FlagCompletion(carapace.ActionMap{
		"branch": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) < 1 {
				return carapace.ActionValues()
			}
			return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: c.Args[0], Branches: true})
		}),
		"reference": carapace.ActionDirectories(),
	})

	carapace.Gen(submodule_addCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionDirectories(),
	)
}
