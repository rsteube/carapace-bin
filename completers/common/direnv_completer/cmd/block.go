package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/direnv"
	"github.com/spf13/cobra"
)

var blockCmd = &cobra.Command{
	Use:     "block",
	Aliases: []string{"deny", "revoke"},
	Short:   "Revokes the authorization of a given .envrc or .env file.",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blockCmd).Standalone()

	rootCmd.AddCommand(blockCmd)

	carapace.Gen(blockCmd).PositionalCompletion(
		direnv.ActionAuths(),
	)
}
