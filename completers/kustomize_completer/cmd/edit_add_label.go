package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_add_labelCmd = &cobra.Command{
	Use:   "label",
	Short: "Adds one or more commonLabels to kustomization.yaml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_add_labelCmd).Standalone()
	edit_add_labelCmd.Flags().BoolP("force", "f", false, "overwrite commonLabel if it already exists")
	edit_addCmd.AddCommand(edit_add_labelCmd)
}
