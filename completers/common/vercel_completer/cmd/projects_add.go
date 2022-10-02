package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var projects_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projects_addCmd).Standalone()

	projectsCmd.AddCommand(projects_addCmd)
}
