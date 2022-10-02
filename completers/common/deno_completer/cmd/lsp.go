package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var lspCmd = &cobra.Command{
	Use:   "lsp",
	Short: "Start the language server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lspCmd).Standalone()

	rootCmd.AddCommand(lspCmd)
}
