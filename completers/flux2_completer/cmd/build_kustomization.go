package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var build_kustomizationCmd = &cobra.Command{
	Use:   "kustomization",
	Short: "Build Kustomization",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_kustomizationCmd).Standalone()
	build_kustomizationCmd.Flags().String("path", "", "Path to the manifests location.)")
	buildCmd.AddCommand(build_kustomizationCmd)
}
