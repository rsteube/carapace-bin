package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_remove_labelCmd = &cobra.Command{
	Use:   "label",
	Short: "Removes one or more commonLabels from kustomization.yaml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_remove_labelCmd).Standalone()
	edit_remove_labelCmd.Flags().BoolP("ignore-non-existence", "i", false, "ignore error if the given label doesn't exist")
	edit_removeCmd.AddCommand(edit_remove_labelCmd)
}
