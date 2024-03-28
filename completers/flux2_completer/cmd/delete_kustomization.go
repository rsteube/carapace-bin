package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var delete_kustomizationCmd = &cobra.Command{
	Use:   "kustomization",
	Short: "Delete a Kustomization resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delete_kustomizationCmd).Standalone()
	deleteCmd.AddCommand(delete_kustomizationCmd)
}
