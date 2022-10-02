package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze Dart code in a directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(analyzeCmd).Standalone()

	analyzeCmd.Flags().Bool("fatal-infos", false, "Treat info level issues as fatal.")
	analyzeCmd.Flags().Bool("fatal-warnings", false, "Treat warning level issues as fatal.")
	analyzeCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	analyzeCmd.Flags().Bool("no-fatal-warnings", false, "Do not treat warning level issues as fatal.")
	rootCmd.AddCommand(analyzeCmd)

	carapace.Gen(analyzeCmd).PositionalCompletion(
		carapace.ActionFiles(".dart"),
	)
}
