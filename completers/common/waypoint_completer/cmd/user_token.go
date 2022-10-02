package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var user_tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Request a new token to access the server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(user_tokenCmd).Standalone()

	user_tokenCmd.Flags().String("expires-in", "", "The duration until the token expires.")
	user_tokenCmd.Flags().String("username", "", "Username to generate the login token for.")

	addGlobalOptions(user_tokenCmd)

	userCmd.AddCommand(user_tokenCmd)
}
