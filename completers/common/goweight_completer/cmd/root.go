package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goweight",
	Short: "A tool to analyze and troubleshoot a Go binary size",
	Long:  "https://github.com/jondot/goweight",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "Show context-sensitive help")
	rootCmd.Flags().BoolP("json", "j", false, "Output json")
	rootCmd.Flags().String("tags", "", "Build tags")
	rootCmd.Flags().Bool("version", false, "Show application version")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
