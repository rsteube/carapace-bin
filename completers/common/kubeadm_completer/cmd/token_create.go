package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var token_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create bootstrap tokens on the server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_createCmd).Standalone()
	token_createCmd.Flags().String("certificate-key", "", "When used together with '--print-join-command', print the full 'kubeadm join' flag needed to join the cluster as a control-plane. To create a new certificate key you must use 'kubeadm init phase upload-certs --upload-certs'.")
	token_createCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	token_createCmd.Flags().String("description", "", "A human friendly description of how this token is used.")
	token_createCmd.Flags().StringSlice("groups", []string{"system:bootstrappers:kubeadm:default-node-token"}, "Extra groups that this token will authenticate as when used for authentication. Must match \"\\Asystem:bootstrappers:[a-z0-9:-]{0,255}[a-z0-9]\\z\"")
	token_createCmd.Flags().Bool("print-join-command", false, "Instead of printing only the token, print the full 'kubeadm join' flag needed to join the cluster using the token.")
	token_createCmd.Flags().String("ttl", "", "The duration before the token is automatically deleted (e.g. 1s, 2m, 3h). If set to '0', the token will never expire")
	token_createCmd.Flags().StringSlice("usages", []string{"signing", "authentication"}, "Describes the ways in which this token can be used. You can pass --usages multiple times or provide a comma separated list of options. Valid options: [signing,authentication]")
	tokenCmd.AddCommand(token_createCmd)

	carapace.Gen(token_createCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
		"usages": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("signing", "authentication").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
