package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_set_labelCmd = &cobra.Command{
	Use:   "label",
	Short: "Sets one or more commonLabels in kustomization.yaml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_set_labelCmd).Standalone()
	edit_setCmd.AddCommand(edit_set_labelCmd)
}
