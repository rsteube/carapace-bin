package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var workspace_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the name of the current workspac",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_showCmd).Standalone()

	workspaceCmd.AddCommand(workspace_showCmd)
}
