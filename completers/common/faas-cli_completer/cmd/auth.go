package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Obtain a token for your OpenFaaS gateway",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	authCmd.Flags().String("audience", "", "OAuth2 audience")
	authCmd.Flags().String("auth-url", "", "OAuth2 Authorize URL i.e. http://idp/oauth/authorize")
	authCmd.Flags().String("client-id", "", "OAuth2 client_id")
	authCmd.Flags().String("client-secret", "", "OAuth2 client_secret, for use with client_credentials grant")
	authCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	authCmd.Flags().String("grant", "implicit", "grant for OAuth2 flow - either implicit, implicit-id or client_credentials")
	authCmd.Flags().Bool("launch-browser", true, "Launch browser for OAuth2 redirect")
	authCmd.Flags().Int("listen-port", 31111, "OAuth2 local port for receiving cookie")
	authCmd.Flags().String("redirect-host", "http://127.0.0.1", "Host for OAuth2 redirection in the implicit flow including URL scheme")
	authCmd.Flags().String("scope", "openid profile", "scope for OAuth2 flow - i.e. \"openid profile\"")
	rootCmd.AddCommand(authCmd)

	carapace.Gen(authCmd).FlagCompletion(carapace.ActionMap{
		"grant": carapace.ActionValues("implicit", "implicit-id", "client_credentials"),
	})
}
