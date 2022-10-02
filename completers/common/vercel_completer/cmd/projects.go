package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Manages your Projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projectsCmd).Standalone()

	rootCmd.AddCommand(projectsCmd)
}
