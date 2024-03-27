package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_remove_transformerCmd = &cobra.Command{
	Use:   "transformer",
	Short: "Removes one or more transformers from kustomization.yaml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_remove_transformerCmd).Standalone()
	edit_removeCmd.AddCommand(edit_remove_transformerCmd)
}
