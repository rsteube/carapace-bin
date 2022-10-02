package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var repo_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for GitLab repositories and projects by name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_searchCmd).Standalone()
	repo_searchCmd.Flags().IntP("page", "p", 1, "Page number")
	repo_searchCmd.Flags().IntP("per-page", "P", 20, "Number of items to list per page")
	repo_searchCmd.Flags().StringP("search", "s", "", "A string contained in the project name")
	repoCmd.AddCommand(repo_searchCmd)
}
