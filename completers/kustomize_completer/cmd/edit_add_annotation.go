package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_add_annotationCmd = &cobra.Command{
	Use:   "annotation",
	Short: "Adds one or more commonAnnotations to kustomization.yaml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_add_annotationCmd).Standalone()
	edit_add_annotationCmd.Flags().BoolP("force", "f", false, "overwrite commonAnnotation if it already exists")
	edit_addCmd.AddCommand(edit_add_annotationCmd)
}
