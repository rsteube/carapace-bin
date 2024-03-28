package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var resume_sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Resume sources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resume_sourceCmd).Standalone()
	resumeCmd.AddCommand(resume_sourceCmd)
}
