package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var resume_helmreleaseCmd = &cobra.Command{
	Use:   "helmrelease",
	Short: "Resume a suspended HelmRelease",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resume_helmreleaseCmd).Standalone()
	resumeCmd.AddCommand(resume_helmreleaseCmd)
}
