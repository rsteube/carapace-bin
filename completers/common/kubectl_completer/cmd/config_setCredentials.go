package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_setCredentialsCmd = &cobra.Command{
	Use:   "set-credentials",
	Short: "Sets a user entry in kubeconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCredentialsCmd).Standalone()

	config_setCredentialsCmd.Flags().String("auth-provider", "", "Auth provider for the user entry in kubeconfig")
	config_setCredentialsCmd.Flags().String("auth-provider-arg", "", "'key=value' arguments for the auth provider")
	config_setCredentialsCmd.Flags().String("embed-certs", "", "Embed client cert/key for the user entry in kubeconfig")
	config_setCredentialsCmd.Flags().String("exec-api-version", "", "API version of the exec credential plugin for the user entry in kubeconfig")
	config_setCredentialsCmd.Flags().String("exec-arg", "", "New arguments for the exec credential plugin command for the user entry in kubeconfig")
	config_setCredentialsCmd.Flags().String("exec-command", "", "Command for the exec credential plugin for the user entry in kubeconfig")
	config_setCredentialsCmd.Flags().String("exec-env", "", "'key=value' environment values for the exec credential plugin")
	configCmd.AddCommand(config_setCredentialsCmd)
}
