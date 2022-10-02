package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var issue_boardCmd = &cobra.Command{
	Use:   "board",
	Short: "Work with GitLab Issue Boards in the given project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_boardCmd).Standalone()
	issueCmd.AddCommand(issue_boardCmd)
}
