package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var token_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete bootstrap tokens on the server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_deleteCmd).Standalone()
	tokenCmd.AddCommand(token_deleteCmd)
}
