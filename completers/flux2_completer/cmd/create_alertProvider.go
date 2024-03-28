package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var create_alertProviderCmd = &cobra.Command{
	Use:   "alert-provider",
	Short: "Create or update a Provider resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_alertProviderCmd).Standalone()
	create_alertProviderCmd.Flags().String("address", "", "path to either the git repository, chat provider or webhook")
	create_alertProviderCmd.Flags().String("channel", "", "channel to send messages to in the case of a chat provider")
	create_alertProviderCmd.Flags().String("secret-ref", "", "name of secret containing authentication token")
	create_alertProviderCmd.Flags().String("type", "", "type of provider")
	create_alertProviderCmd.Flags().String("username", "", "bot username used by the provider")
	createCmd.AddCommand(create_alertProviderCmd)
}
