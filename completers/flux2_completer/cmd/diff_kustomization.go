package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var diff_kustomizationCmd = &cobra.Command{
	Use:   "kustomization",
	Short: "Diff Kustomization",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diff_kustomizationCmd).Standalone()
	diff_kustomizationCmd.Flags().String("path", "", "Path to a local directory that matches the specified Kustomization.spec.path.)")
	diffCmd.AddCommand(diff_kustomizationCmd)
}
