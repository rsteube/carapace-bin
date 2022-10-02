package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Apply automated fixes to Dart source code",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fixCmd).Standalone()

	fixCmd.Flags().Bool("apply", false, "Apply the proposed changes.")
	fixCmd.Flags().BoolP("dry-run", "n", false, "Preview the proposed changes but make no changes.")
	fixCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	rootCmd.AddCommand(fixCmd)

	carapace.Gen(fixCmd).PositionalCompletion(
		carapace.ActionFiles(".dart"),
	)
}
