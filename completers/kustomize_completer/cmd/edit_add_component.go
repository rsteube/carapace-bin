package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_add_componentCmd = &cobra.Command{
	Use:   "component",
	Short: "Add the name of a file containing a component to the kustomization file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_add_componentCmd).Standalone()
	edit_addCmd.AddCommand(edit_add_componentCmd)
}
