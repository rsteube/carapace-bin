package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "package a chart directory into a chart archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	packageCmd.Flags().String("app-version", "", "set the appVersion on the chart to this version")
	packageCmd.Flags().BoolP("dependency-update", "u", false, "update dependencies from \"Chart.yaml\" to dir \"charts/\" before packaging")
	packageCmd.Flags().StringP("destination", "d", ".", "location to write the chart.")
	packageCmd.Flags().String("key", "", "name of the key to use when signing. Used if --sign is true")
	packageCmd.Flags().String("keyring", "/home/user/.gnupg/pubring.gpg", "location of a public keyring")
	packageCmd.Flags().String("passphrase-file", "", "location of a file which contains the passphrase for the signing key. Use \"-\" in order to read from stdin.")
	packageCmd.Flags().Bool("sign", false, "use a PGP private key to sign this package")
	packageCmd.Flags().String("version", "", "set the version on the chart to this semver version")
	rootCmd.AddCommand(packageCmd)

	carapace.Gen(packageCmd).FlagCompletion(carapace.ActionMap{
		"destination":     carapace.ActionDirectories(),
		"keyring":         carapace.ActionFiles(),
		"passphrase-file": carapace.ActionFiles(),
	})

	carapace.Gen(packageCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
