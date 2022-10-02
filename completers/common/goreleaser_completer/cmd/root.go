package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goreleaser",
	Short: "Deliver Go binaries as fast and easily as possible",
	Long:  "https://goreleaser.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("debug", false, "Enable debug mode")
	rootCmd.Flags().BoolP("help", "h", false, "help for goreleaser")
	rootCmd.Flags().BoolP("version", "v", false, "version for goreleaser")
}
