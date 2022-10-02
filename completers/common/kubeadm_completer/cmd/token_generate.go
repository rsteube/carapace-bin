package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var token_generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate and print a bootstrap token, but do not create it on the server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_generateCmd).Standalone()
	tokenCmd.AddCommand(token_generateCmd)
}
