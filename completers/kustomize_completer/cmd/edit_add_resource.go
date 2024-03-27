package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_add_resourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "Add the name of a file containing a resource to the kustomization file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_add_resourceCmd).Standalone()
	edit_addCmd.AddCommand(edit_add_resourceCmd)
}
