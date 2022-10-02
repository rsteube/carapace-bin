package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var token_revokeCmd = &cobra.Command{
	Use:   "revoke",
	Short: "remove an authentication token from the registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_revokeCmd).Standalone()

	tokenCmd.AddCommand(token_revokeCmd)
}
