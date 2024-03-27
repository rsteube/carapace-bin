package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_remove_resourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "Removes one or more resource file paths from kustomization.yaml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_remove_resourceCmd).Standalone()
	edit_removeCmd.AddCommand(edit_remove_resourceCmd)
}
