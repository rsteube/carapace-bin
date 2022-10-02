package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var repo_deployKey_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List deploy keys in a GitHub repository",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_deployKey_listCmd).Standalone()
	repo_deployKeyCmd.AddCommand(repo_deployKey_listCmd)
}
