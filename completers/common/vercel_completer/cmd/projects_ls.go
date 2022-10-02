package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var projects_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Show all projects in the selected team/user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projects_lsCmd).Standalone()

	projectsCmd.AddCommand(projects_lsCmd)
}
