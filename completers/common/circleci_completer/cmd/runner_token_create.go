package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runner_token_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a token for a resource-class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_token_createCmd).Standalone()
	runner_tokenCmd.AddCommand(runner_token_createCmd)
}
