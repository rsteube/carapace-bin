package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var resume_source_chartCmd = &cobra.Command{
	Use:   "chart",
	Short: "Resume a suspended HelmChart",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resume_source_chartCmd).Standalone()
	resume_sourceCmd.AddCommand(resume_source_chartCmd)
}
