package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var resumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Resume suspended resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resumeCmd).Standalone()
	resumeCmd.PersistentFlags().Bool("all", false, "resume all resources in that namespace")
	rootCmd.AddCommand(resumeCmd)
}
