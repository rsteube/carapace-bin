package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runner_token_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a token",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_token_deleteCmd).Standalone()
	runner_tokenCmd.AddCommand(runner_token_deleteCmd)
}
