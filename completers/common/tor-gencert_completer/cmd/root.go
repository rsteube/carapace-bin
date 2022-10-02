package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tor-gencert",
	Short: "Generate certs and keys for Tor directory authorities",
	Long:  "https://www.torproject.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("a", "a", "", "advertise the address:port combination as this authority’s preferred directory port")
	rootCmd.Flags().StringS("c", "c", "", "Write the certificate to the specified file")
	rootCmd.Flags().Bool("create-identity-key", false, "Generate a new identity key")
	rootCmd.Flags().BoolP("help", "h", false, "Display help text and exit")
	rootCmd.Flags().StringS("i", "i", "", "Read the identity key from the specified file")
	rootCmd.Flags().StringS("m", "m", "", "Number of months that the certificate should be valid")
	rootCmd.Flags().String("passphrase-fd", "", "Filedescriptor to read the passphrase from")
	rootCmd.Flags().BoolP("reuse", "r", false, "Generate a new certificate, but not a new signing key")
	rootCmd.Flags().StringS("s", "s", "", "Write the signing key to the specified file")
	rootCmd.Flags().BoolS("v", "v", false, "Display verbose output")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"c": carapace.ActionFiles(),
		"i": carapace.ActionFiles(),
		"s": carapace.ActionFiles(),
	})
}
