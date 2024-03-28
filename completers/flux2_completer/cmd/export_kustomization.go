package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var export_kustomizationCmd = &cobra.Command{
	Use:   "kustomization",
	Short: "Export Kustomization resources in YAML format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(export_kustomizationCmd).Standalone()
	exportCmd.AddCommand(export_kustomizationCmd)
}
