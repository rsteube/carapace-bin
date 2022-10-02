package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var token_listCmd = &cobra.Command{
	Use:   "list",
	Short: "show a table of all active authentication tokens",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_listCmd).Standalone()

	tokenCmd.AddCommand(token_listCmd)
}
