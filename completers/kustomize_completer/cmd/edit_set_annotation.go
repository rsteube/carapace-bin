package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_set_annotationCmd = &cobra.Command{
	Use:   "annotation",
	Short: "Sets one or more commonAnnotations in kustomization.yaml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_set_annotationCmd).Standalone()
	edit_setCmd.AddCommand(edit_set_annotationCmd)
}
