package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var resume_source_helmCmd = &cobra.Command{
	Use:   "helm",
	Short: "Resume a suspended HelmRepository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resume_source_helmCmd).Standalone()
	resume_sourceCmd.AddCommand(resume_source_helmCmd)
}
