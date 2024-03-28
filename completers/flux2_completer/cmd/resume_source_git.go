package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var resume_source_gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Resume a suspended GitRepository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resume_source_gitCmd).Standalone()
	resume_sourceCmd.AddCommand(resume_source_gitCmd)
}
