package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var gpgKey_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a GPG key to your GitHub account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gpgKey_addCmd).Standalone()
	gpgKeyCmd.AddCommand(gpgKey_addCmd)

	carapace.Gen(gpgKey_addCmd).PositionalCompletion(
		carapace.ActionFiles(".pgp"),
	)
}
