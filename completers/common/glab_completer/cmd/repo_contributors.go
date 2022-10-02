package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_contributorsCmd = &cobra.Command{
	Use:   "contributors",
	Short: "Get repository contributors list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_contributorsCmd).Standalone()
	repo_contributorsCmd.Flags().StringP("order", "o", "commits", "Return contributors ordered by name, email, or commits (orders by commit date) fields")
	repo_contributorsCmd.Flags().IntP("page", "p", 1, "Page number")
	repo_contributorsCmd.Flags().IntP("per-page", "P", 30, "Number of items to list per page.")
	repo_contributorsCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format or full URL or git URL")
	repo_contributorsCmd.Flags().StringP("sort", "s", "", "Return contributors sorted in asc or desc order")
	repoCmd.AddCommand(repo_contributorsCmd)

	carapace.Gen(repo_contributorsCmd).FlagCompletion(carapace.ActionMap{
		"order": carapace.ActionValues("name", "email", "commits"),
		"repo":  action.ActionRepo(repo_contributorsCmd),
		"sort":  carapace.ActionValues("asc", "desc"),
	})
}
