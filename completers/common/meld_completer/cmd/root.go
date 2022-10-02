package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "meld",
	Short: "Meld is a file and directory comparison tool",
	Long:  "https://meldmerge.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("auto-compare", "a", false, "Automatically compare all differing files on startup")
	rootCmd.Flags().Bool("auto-merge", false, "Automatically merge files")
	rootCmd.Flags().String("comparison-file", "", "Load a saved comparison from a Meld comparison file")
	rootCmd.Flags().Bool("diff", false, "Create a diff tab for the supplied files or folders")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().StringP("label", "l", "", "Set label to use instead of file name")
	rootCmd.Flags().BoolP("newtab", "n", false, "Open a new tab in an already running instance")
	rootCmd.Flags().StringP("output", "o", "", "Set the target file for saving a merge result")
	rootCmd.Flags().BoolP("unified", "u", false, "Ignored for compatibility")
	rootCmd.Flags().Bool("version", false, "show program's version number and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"comparison-file": carapace.ActionFiles(),
		"output":          carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
