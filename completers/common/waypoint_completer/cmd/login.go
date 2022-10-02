package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/waypoint_completer/cmd/action"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to a Waypoint server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginCmd).Standalone()

	loginCmd.Flags().String("auth-method", "", "Auth method to use for login.")
	loginCmd.Flags().Bool("from-kubernetes", false, "Perform the initial authentication after Waypoint is installed to Kubernetes.")
	loginCmd.Flags().String("from-kubernetes-namespace", "", "The name of the Kubernetes namespace that has the Waypoint token.")
	loginCmd.Flags().String("from-kubernetes-secret", "", "The name of the Kubernetes secret that has the Waypoint token.")
	loginCmd.Flags().String("from-kubernetes-service", "", "The name of the Kubernetes service to get the server address from.")
	loginCmd.Flags().String("token", "", "Auth with a token.")

	addGlobalOptions(loginCmd)
	addConnectionOptions(loginCmd)

	rootCmd.AddCommand(loginCmd)

	carapace.Gen(loginCmd).FlagCompletion(carapace.ActionMap{
		"auth-method": action.ActionAuthMethods(),
	})
}
