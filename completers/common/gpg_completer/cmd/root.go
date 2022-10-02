package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gpg",
	Short: "OpenPGP encryption and signing tool",
	Long:  "https://gnupg.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("armor", "a", false, "create ascii armored output")
	rootCmd.Flags().Bool("card-status", false, "print the card status")
	rootCmd.Flags().Bool("change-passphrase", false, "change a passphrase")
	rootCmd.Flags().Bool("change-pin", false, "change a card's PIN")
	rootCmd.Flags().Bool("check-signatures", false, "list and check key signatures")
	rootCmd.Flags().Bool("clear-sign", false, "make a clear text signature")
	rootCmd.Flags().BoolP("decrypt", "d", false, "decrypt data (default)")
	rootCmd.Flags().Bool("delete-keys", false, "remove keys from the public keyring")
	rootCmd.Flags().Bool("delete-secret-keys", false, "remove keys from the secret keyring")
	rootCmd.Flags().BoolP("detach-sign", "b", false, "make a detached signature")
	rootCmd.Flags().BoolP("dry-run", "n", false, "do not make any changes")
	rootCmd.Flags().Bool("edit-card", false, "change data on a card")
	rootCmd.Flags().Bool("edit-key", false, "sign or edit a key")
	rootCmd.Flags().BoolP("encrypt", "e", false, "encrypt data")
	rootCmd.Flags().Bool("export", false, "export keys")
	rootCmd.Flags().Bool("fingerprint", false, "list keys and fingerprints")
	rootCmd.Flags().Bool("full-generate-key", false, "full featured key pair generation")
	rootCmd.Flags().Bool("generate-key", false, "generate a new key pair")
	rootCmd.Flags().Bool("generate-revocation", false, "generate a revocation certificate")
	rootCmd.Flags().Bool("import", false, "import/merge keys")
	rootCmd.Flags().BoolP("interactive", "i", false, "prompt before overwriting")
	rootCmd.Flags().BoolP("list-keys", "k", false, "list keys")
	rootCmd.Flags().BoolP("list-secret-keys", "K", false, "list secret keys")
	rootCmd.Flags().Bool("list-signatures", false, "list keys and signatures")
	rootCmd.Flags().StringP("local-user", "u", "", "use USER-ID to sign or decrypt")
	rootCmd.Flags().Bool("lsign-key", false, "sign a key locally")
	rootCmd.Flags().Bool("openpgp", false, "use strict OpenPGP behavior")
	rootCmd.Flags().StringP("output", "o", "", "write output to FILE")
	rootCmd.Flags().Bool("print-md", false, "print message digests")
	rootCmd.Flags().Bool("quick-add-uid", false, "quickly add a new user-id")
	rootCmd.Flags().Bool("quick-generate-key", false, "quickly generate a new key pair")
	rootCmd.Flags().Bool("quick-lsign-key", false, "quickly sign a key locally")
	rootCmd.Flags().Bool("quick-revoke-sig", false, "quickly revoke a key signature")
	rootCmd.Flags().Bool("quick-revoke-uid", false, "quickly revoke a user-id")
	rootCmd.Flags().Bool("quick-set-expire", false, "quickly set a new expiration date")
	rootCmd.Flags().Bool("quick-sign-key", false, "quickly sign a key")
	rootCmd.Flags().Bool("receive-keys", false, "import keys from a keyserver")
	rootCmd.Flags().StringP("recipient", "r", "", "encrypt for USER-ID")
	rootCmd.Flags().Bool("refresh-keys", false, "update all keys from a keyserver")
	rootCmd.Flags().Bool("search-keys", false, "search for keys on a keyserver")
	rootCmd.Flags().Bool("send-keys", false, "export keys to a keyserver")
	rootCmd.Flags().Bool("server", false, "run in server mode")
	rootCmd.Flags().BoolP("sign", "s", false, "make a signature")
	rootCmd.Flags().Bool("sign-key", false, "sign a key")
	rootCmd.Flags().BoolP("symmetric", "c", false, "encryption only with symmetric cipher")
	rootCmd.Flags().Bool("textmode", false, "use canonical text mode")
	rootCmd.Flags().String("tofu-policy", "", "set the TOFU policy for a key")
	rootCmd.Flags().Bool("update-trustdb", false, "update the trust database")
	rootCmd.Flags().BoolP("verbose", "v", false, "verbose")
	rootCmd.Flags().Bool("verify", false, "verify a signature")
	rootCmd.Flags().StringS("z", "z", "", "set compress level to N (0 disables)")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"local-user":  os.ActionUsers(),
		"output":      carapace.ActionFiles(),
		"recipient":   os.ActionUsers(),
		"tofu-policy": carapace.ActionValues("auto", "good", "unknown", "bad", "ask"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
