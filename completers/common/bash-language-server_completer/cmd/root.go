package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bash-language-server",
	Short: "A language server for Bash",
	Long:  "https://github.com/bash-lsp/bash-language-server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().BoolP("version", "v", false, "show version")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("start"),
	)
}
