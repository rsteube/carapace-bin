package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serveCmd).Standalone()
	serveCmd.Flags().StringP("dir", "d", "", "indicate a public directory path")
	serveCmd.Flags().BoolP("help", "h", false, "display help for command")
	serveCmd.Flags().StringP("port", "p", "", "port used (default: 4000)")
	serveCmd.Flags().StringP("resume", "r", "", "resume file")
	serveCmd.Flags().BoolP("silent", "s", false, "open browser")
	serveCmd.Flags().StringP("theme", "t", "", "theme name")

	rootCmd.AddCommand(serveCmd)

	carapace.Gen(serveCmd).FlagCompletion(carapace.ActionMap{
		"dir":    carapace.ActionDirectories(),
		"resume": carapace.ActionFiles(),
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
}
