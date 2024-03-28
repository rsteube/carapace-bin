package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var resume_kustomizationCmd = &cobra.Command{
	Use:   "kustomization",
	Short: "Resume a suspended Kustomization",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resume_kustomizationCmd).Standalone()
	resumeCmd.AddCommand(resume_kustomizationCmd)
}
