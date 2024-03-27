package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets the value of different fields in kustomization file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_setCmd).Standalone()
	editCmd.AddCommand(edit_setCmd)
}
