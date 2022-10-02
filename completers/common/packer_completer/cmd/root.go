package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "packer",
	Short: "Create identical machine images for multiple platforms from a single source configuration.",
	Long:  "https://www.packer.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolP("version", "v", false, "show version")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "show help")
}
