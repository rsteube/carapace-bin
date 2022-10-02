package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "volta",
	Short: "The JavaScript Launcher",
	Long:  "https://volta.sh/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.Flags().Bool("quiet", false, "Prevents unnecessary output")
	rootCmd.Flags().Bool("verbose", false, "Enables verbose diagnostics")
	rootCmd.Flags().BoolP("version", "v", false, "Prints the current version of Volta")
}
