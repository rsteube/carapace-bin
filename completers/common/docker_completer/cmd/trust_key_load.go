package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var trust_key_loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load a private key file for signing",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trust_key_loadCmd).Standalone()

	trust_key_loadCmd.Flags().String("name", "", "Name for the loaded key (default \"signer\")")
	trust_keyCmd.AddCommand(trust_key_loadCmd)

	carapace.Gen(trust_key_loadCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
