package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "turbo",
	Short: "The High-performance Build System for JavaScript & TypeScript Codebases",
	Long:  "https://turborepo.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().Bool("color", false, "forces the use of color")
	rootCmd.PersistentFlags().Bool("no-color", false, "suppresses the use of color")
	rootCmd.Flags().BoolP("version", "v", false, "show version")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "show help")
}
