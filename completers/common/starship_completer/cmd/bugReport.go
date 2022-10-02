package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bugReportCmd = &cobra.Command{
	Use:   "bug-report",
	Short: "Create a pre-populated GitHub issue with information about your configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bugReportCmd).Standalone()

	bugReportCmd.Flags().BoolP("help", "h", false, "Prints help information")
	bugReportCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.AddCommand(bugReportCmd)
}
