package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var dependency_buildCmd = &cobra.Command{
	Use:   "build",
	Short: "rebuild the charts/ directory based on the Chart.lock file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	dependency_buildCmd.Flags().String("keyring", "/home/user/.gnupg/pubring.gpg", "keyring containing public keys")
	dependency_buildCmd.Flags().Bool("skip-refresh", false, "do not refresh the local repository cache")
	dependency_buildCmd.Flags().Bool("verify", false, "verify the packages against signatures")
	dependencyCmd.AddCommand(dependency_buildCmd)

	carapace.Gen(dependency_buildCmd).FlagCompletion(carapace.ActionMap{
		"keyring": carapace.ActionFiles(),
	})

	carapace.Gen(dependency_buildCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
