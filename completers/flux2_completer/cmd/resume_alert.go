package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var resume_alertCmd = &cobra.Command{
	Use:   "alert",
	Short: "Resume a suspended Alert",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resume_alertCmd).Standalone()
	resumeCmd.AddCommand(resume_alertCmd)
}
