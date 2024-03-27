package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_set_nameprefixCmd = &cobra.Command{
	Use:   "nameprefix",
	Short: "Sets the value of the namePrefix field in the kustomization file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_set_nameprefixCmd).Standalone()
	edit_setCmd.AddCommand(edit_set_nameprefixCmd)
}
