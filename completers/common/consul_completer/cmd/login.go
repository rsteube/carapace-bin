package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Consul using an auth method",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginCmd).Standalone()
	addClientFlags(loginCmd)
	addServerFlags(loginCmd)

	loginCmd.Flags().String("bearer-token-file", "", "Path to a file containing a secret bearer token to use with this auth method.")
	loginCmd.Flags().StringArray("meta", []string{}, "Metadata to set on the token, formatted as key=value.")
	loginCmd.Flags().String("method", "", "Name of the auth method to login to.")
	loginCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	loginCmd.Flags().String("token-sink-file", "", "The most recent token's SecretID is kept up to date in this file.")
	loginCmd.Flags().String("type", "", "Type of the auth method to login to.")
	rootCmd.AddCommand(loginCmd)

	carapace.Gen(loginCmd).FlagCompletion(carapace.ActionMap{
		"bearer-token-file": carapace.ActionFiles(),
		"token-sink-file":   carapace.ActionFiles(),
	})
}
