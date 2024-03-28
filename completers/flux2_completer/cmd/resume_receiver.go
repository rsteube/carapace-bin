package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var resume_receiverCmd = &cobra.Command{
	Use:   "receiver",
	Short: "Resume a suspended Receiver",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resume_receiverCmd).Standalone()
	resumeCmd.AddCommand(resume_receiverCmd)
}
