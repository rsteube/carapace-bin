package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var diagnosticCmd = &cobra.Command{
	Use:   "diagnostic",
	Short: "Check the status of your CircleCI CLI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diagnosticCmd).Standalone()
	rootCmd.AddCommand(diagnosticCmd)
}
