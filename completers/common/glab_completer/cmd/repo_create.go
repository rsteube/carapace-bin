package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new GitLab project/repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_createCmd).Standalone()
	repo_createCmd.Flags().String("defaultBranch", "", "Default branch of the project. If not provided, `master` by default.")
	repo_createCmd.Flags().StringP("description", "d", "", "Description of the new project")
	repo_createCmd.Flags().StringP("group", "g", "", "Namespace/group for the new project (defaults to the current user’s namespace)")
	repo_createCmd.Flags().Bool("internal", false, "Make project internal: visible to any authenticated user (default)")
	repo_createCmd.Flags().StringP("name", "n", "", "Name of the new project")
	repo_createCmd.Flags().BoolP("private", "p", false, "Make project private: visible only to project members")
	repo_createCmd.Flags().BoolP("public", "P", false, "Make project public: visible without any authentication")
	repo_createCmd.Flags().Bool("readme", false, "Initialize project with README.md")
	repo_createCmd.Flags().String("remoteName", "origin", "Remote name for the Git repository you're in. If not provided, `origin` by default.")
	repo_createCmd.Flags().StringArrayP("tag", "t", []string{}, "The list of tags for the project.")
	repoCmd.AddCommand(repo_createCmd)

	carapace.Gen(repo_createCmd).FlagCompletion(carapace.ActionMap{
		"group": action.ActionGroups(repo_createCmd),
	})

	carapace.Gen(repo_createCmd).PositionalCompletion(
		action.ActionRepo(repo_createCmd),
	)
}
