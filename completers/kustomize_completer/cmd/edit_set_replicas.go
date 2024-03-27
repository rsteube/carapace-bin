package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_set_replicasCmd = &cobra.Command{
	Use:   "replicas",
	Short: "Sets replicas count for resources in the kustomization file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_set_replicasCmd).Standalone()
	edit_setCmd.AddCommand(edit_set_replicasCmd)
}
