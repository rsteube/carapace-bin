package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds an item to the kustomization file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_addCmd).Standalone()
	editCmd.AddCommand(edit_addCmd)
}
