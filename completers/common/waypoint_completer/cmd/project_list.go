package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var project_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all registered projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_listCmd).Standalone()

	projectCmd.AddCommand(project_listCmd)
}
