package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes items from the kustomization file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_removeCmd).Standalone()
	editCmd.AddCommand(edit_removeCmd)
}
