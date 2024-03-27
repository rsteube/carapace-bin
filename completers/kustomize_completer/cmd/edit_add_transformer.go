package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_add_transformerCmd = &cobra.Command{
	Use:   "transformer",
	Short: "Add the name of a file containing a transformer configuration to the kustomization file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_add_transformerCmd).Standalone()
	edit_addCmd.AddCommand(edit_add_transformerCmd)
}
