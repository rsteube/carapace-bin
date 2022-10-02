package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clamsubmit",
	Short: "File submission utility for ClamAV",
	Long:  "http://www.clamav.net/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("?", "?", false, "Show this help")
	rootCmd.Flags().StringS("N", "N", "", "Your name contained in quotation marks (required)")
	rootCmd.Flags().StringS("V", "V", "", "Detected virus name (required with -p)")
	rootCmd.Flags().BoolS("d", "d", false, "Enable debug output")
	rootCmd.Flags().StringS("e", "e", "", "Your email address (required)")
	rootCmd.Flags().BoolS("h", "h", false, "Show this help")
	rootCmd.Flags().StringS("n", "n", "", "Submit a false negative (FN)")
	rootCmd.Flags().StringS("p", "p", "", "Submit a false positive (FP)")
	rootCmd.Flags().BoolS("v", "v", false, "Show version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"n": carapace.ActionFiles(),
		"p": carapace.ActionFiles(),
	})
}
