package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_set_namespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "Sets the value of the namespace field in the kustomization file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_set_namespaceCmd).Standalone()
	edit_setCmd.AddCommand(edit_set_namespaceCmd)
}
