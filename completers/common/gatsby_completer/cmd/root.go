package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gatsby",
	Short: "Build blazing fast, modern apps and websites with React",
	Long:  "https://www.gatsbyjs.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Show help")
	rootCmd.Flags().Bool("json", false, "Turn on the JSON logger")
	rootCmd.Flags().Bool("no-color", false, "Turn off the color in output")
	rootCmd.Flags().Bool("no-colors", false, "Turn off the color in output")
	rootCmd.Flags().Bool("verbose", false, "Turn on verbose output")
	rootCmd.Flags().BoolP("version", "v", false, "Show the version of the Gatsby CLI")
}
