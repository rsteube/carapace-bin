package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runner_token_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tokens for a resource-class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_token_listCmd).Standalone()
	runner_tokenCmd.AddCommand(runner_token_listCmd)
}
