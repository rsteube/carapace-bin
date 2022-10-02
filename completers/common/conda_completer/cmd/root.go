package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "conda",
	Short: "conda is a tool for managing and deploying applications, environments and packages",
	Long:  "https://docs.conda.io/en/latest/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Show this help message and exit.")
	rootCmd.Flags().BoolP("version", "V", false, "Show the conda version number and exit.")
}
