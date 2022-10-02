package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubeconfig_userCmd = &cobra.Command{
	Use:   "user",
	Short: "Output a kubeconfig file for an additional user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubeconfig_userCmd).Standalone()
	kubeconfig_userCmd.Flags().String("client-name", "", "The name of user. It will be used as the CN if client certificates are created")
	kubeconfig_userCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	kubeconfig_userCmd.Flags().StringSlice("org", []string{}, "The orgnizations of the client certificate. It will be used as the O if client certificates are created")
	kubeconfig_userCmd.Flags().String("token", "", "The token that should be used as the authentication mechanism for this kubeconfig, instead of client certificates")
	kubeconfig_userCmd.Flags().String("validity-period", "", "The validity period of the client certificate. It is an offset from the current time.")
	kubeconfigCmd.AddCommand(kubeconfig_userCmd)

	carapace.Gen(kubeconfig_userCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
