package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exportCmd).Standalone()
	exportCmd.Flags().BoolP("force", "F", false, "bypass schema testing")
	exportCmd.Flags().StringP("format", "f", "", "output format")
	exportCmd.Flags().BoolP("help", "h", false, "display help for command")
	exportCmd.Flags().StringP("theme", "t", "", "theme name")

	rootCmd.AddCommand(exportCmd)

	carapace.Gen(exportCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("html", "pdf"),
		"theme": carapace.ActionValues(
			"jsonresume-theme-class",
			"jsonresume-theme-classy",
			"jsonresume-theme-elegant",
			"jsonresume-theme-flat",
			"jsonresume-theme-kendall",
			"jsonresume-theme-kwan",
			"jsonresume-theme-md",
			"jsonresume-theme-modern",
			"jsonresume-theme-onepage",
			"jsonresume-theme-paper",
			"jsonresume-theme-short",
			"jsonresume-theme-slick",
			"jsonresume-theme-spartan",
			"jsonresume-theme-stackoverflow",
		),
	})

	carapace.Gen(exportCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
