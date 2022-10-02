package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze the project's Dart code",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(analyzeCmd).Standalone()

	analyzeCmd.Flags().Bool("congratulate", false, "Show output even when there are no errors, warnings, hints, or lints.")
	analyzeCmd.Flags().Bool("current-package", false, "Analyze the current project, if applicable.")
	analyzeCmd.Flags().Bool("fatal-infos", false, "Treat info level issues as fatal.")
	analyzeCmd.Flags().Bool("fatal-warnings", false, "Treat warning level issues as fatal.")
	analyzeCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	analyzeCmd.Flags().Bool("no-congratulate", false, "Do not show output even when there are no errors, warnings, hints, or lints.")
	analyzeCmd.Flags().Bool("no-current-package", false, "Do not analyze the current project, if applicable.")
	analyzeCmd.Flags().Bool("no-fatal-infos", false, "Do not treat info level issues as fatal.")
	analyzeCmd.Flags().Bool("no-fatal-warnings", false, "Do not treat warning level issues as fatal.")
	analyzeCmd.Flags().Bool("no-preamble", false, "When analyzing the flutter repository, do not display the number of files that will be analyzed.")
	analyzeCmd.Flags().Bool("no-pub", false, "Do not run \"flutter pub get\" before executing this command.")
	analyzeCmd.Flags().Bool("preamble", false, "When analyzing the flutter repository, display the number of files that will be analyzed.")
	analyzeCmd.Flags().Bool("pub", false, "Run \"flutter pub get\" before executing this command.")
	analyzeCmd.Flags().Bool("watch", false, "Run analysis continuously, watching the filesystem for changes.")
	analyzeCmd.Flags().String("write", "", "Also output the results to a file.")
	rootCmd.AddCommand(analyzeCmd)

	carapace.Gen(analyzeCmd).FlagCompletion(carapace.ActionMap{
		"write": carapace.ActionFiles(),
	})
}
