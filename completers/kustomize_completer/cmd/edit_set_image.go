package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_set_imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Sets images and their new names, new tags or digests in the kustomization file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_set_imageCmd).Standalone()
	edit_setCmd.AddCommand(edit_set_imageCmd)
}
