package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var repo_createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a new repository",
	Aliases: []string{"new"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_createCmd).Standalone()
	repo_createCmd.Flags().Bool("add-readme", false, "Add a README file to the new repository")
	repo_createCmd.Flags().BoolP("clone", "c", false, "Clone the new repository to the current directory")
	repo_createCmd.Flags().StringP("description", "d", "", "Description of the repository")
	repo_createCmd.Flags().Bool("disable-issues", false, "Disable issues in the new repository")
	repo_createCmd.Flags().Bool("disable-wiki", false, "Disable wiki in the new repository")
	repo_createCmd.Flags().StringP("gitignore", "g", "", "Specify a gitignore template for the repository")
	repo_createCmd.Flags().StringP("homepage", "h", "", "Repository home page `URL`")
	repo_createCmd.Flags().Bool("include-all-branches", false, "Include all branches from template repository")
	repo_createCmd.Flags().Bool("internal", false, "Make the new repository internal")
	repo_createCmd.Flags().StringP("license", "l", "", "Specify an Open Source License for the repository")
	repo_createCmd.Flags().Bool("private", false, "Make the new repository private")
	repo_createCmd.Flags().Bool("public", false, "Make the new repository public")
	repo_createCmd.Flags().Bool("push", false, "Push local commits to the new repository")
	repo_createCmd.Flags().StringP("remote", "r", "", "Specify remote name for the new repository")
	repo_createCmd.Flags().StringP("source", "s", "", "Specify path to local repository to use as source")
	repo_createCmd.Flags().StringP("team", "t", "", "The `name` of the organization team to be granted access")
	repo_createCmd.Flags().StringP("template", "p", "", "Make the new repository based on a template `repository`")
	repoCmd.AddCommand(repo_createCmd)

	carapace.Gen(repo_createCmd).FlagCompletion(carapace.ActionMap{
		"gitignore": action.ActionGitignoreTemplates(repo_createCmd),
		"license":   gh.ActionLicenses(gh.HostOpts{}),
		"source":    carapace.ActionDirectories(),
		// TODO team
		"template": action.ActionOwnerRepositories(repo_createCmd),
	})
}
