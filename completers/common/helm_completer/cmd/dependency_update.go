package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var dependency_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update charts/ based on the contents of Chart.yaml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	dependency_updateCmd.Flags().String("keyring", "/home/user/.gnupg/pubring.gpg", "keyring containing public keys")
	dependency_updateCmd.Flags().Bool("skip-refresh", false, "do not refresh the local repository cache")
	dependency_updateCmd.Flags().Bool("verify", false, "verify the packages against signatures")
	dependencyCmd.AddCommand(dependency_updateCmd)

	carapace.Gen(dependency_updateCmd).FlagCompletion(carapace.ActionMap{
		"keyring": carapace.ActionFiles(),
	})

	carapace.Gen(dependency_updateCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
