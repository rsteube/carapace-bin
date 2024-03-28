package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var resume_imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Resume image automation objects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resume_imageCmd).Standalone()
	resumeCmd.AddCommand(resume_imageCmd)
}
