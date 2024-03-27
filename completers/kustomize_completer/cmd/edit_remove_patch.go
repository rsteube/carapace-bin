package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_remove_patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Removes a patch from kustomization.yaml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_remove_patchCmd).Standalone()
	edit_remove_patchCmd.Flags().String("annotation-selector", "", "annotationSelector in patch target")
	edit_remove_patchCmd.Flags().String("group", "", "API group in patch target")
	edit_remove_patchCmd.Flags().String("kind", "", "Resource kind in patch target")
	edit_remove_patchCmd.Flags().String("label-selector", "", "labelSelector in patch target")
	edit_remove_patchCmd.Flags().String("name", "", "Resource name in patch target")
	edit_remove_patchCmd.Flags().String("namespace", "", "Resource namespace in patch target")
	edit_remove_patchCmd.Flags().String("patch", "", "Literal string of patch content. Cannot be used with --path at the same time.")
	edit_remove_patchCmd.Flags().String("path", "", "Path to the patch file. Cannot be used with --patch at the same time.")
	edit_remove_patchCmd.Flags().String("version", "", "API version in patch target")
	edit_removeCmd.AddCommand(edit_remove_patchCmd)
}
