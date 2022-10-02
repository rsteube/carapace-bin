package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Work with GitLab repositories and projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoCmd).Standalone()
	rootCmd.AddCommand(repoCmd)
}
