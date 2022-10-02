package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "resume-cli",
	Short: "command line tool for JSON Resume",
	Long:  "https://github.com/jsonresume/resume-cli",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolP("help", "h", false, "display help for command")
	rootCmd.Flags().BoolP("version", "V", false, "output the version number")
}
