package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var acl_authMethod_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an ACL auth method",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_authMethod_updateCmd).Standalone()
	addClientFlags(acl_authMethod_updateCmd)
	addServerFlags(acl_authMethod_updateCmd)

	acl_authMethod_updateCmd.Flags().String("config", "", "The configuration for the auth method.")
	acl_authMethod_updateCmd.Flags().String("description", "", "A description of the auth method.")
	acl_authMethod_updateCmd.Flags().String("display-name", "", "An optional name to use instead of the name when displaying this auth method in a UI.")
	acl_authMethod_updateCmd.Flags().String("format", "", "Output format.")
	acl_authMethod_updateCmd.Flags().String("kubernetes-ca-cert", "", "PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.")
	acl_authMethod_updateCmd.Flags().String("kubernetes-host", "", "Address of the Kubernetes API server.")
	acl_authMethod_updateCmd.Flags().String("kubernetes-service-account-jwt", "", "A Kubernetes service account JWT used to access the TokenReview API to validate other JWTs during login.")
	acl_authMethod_updateCmd.Flags().String("max-token-ttl", "", "Duration of time all tokens created by this auth method should be valid for")
	acl_authMethod_updateCmd.Flags().Bool("meta", false, "Indicates that auth method metadata such as the raft indices should be shown for each entry.")
	acl_authMethod_updateCmd.Flags().String("name", "", "The auth method name.")
	acl_authMethod_updateCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_authMethod_updateCmd.Flags().Bool("no-merge", false, "Do not merge the current auth method information with what is provided to the command.")
	acl_authMethod_updateCmd.Flags().String("token-locality", "", "Defines the kind of token that this auth method should produce.")
	acl_authMethodCmd.AddCommand(acl_authMethod_updateCmd)

	carapace.Gen(acl_authMethod_updateCmd).FlagCompletion(carapace.ActionMap{
		// TODO flag completions
		"config":             action.ActionOptionalFiles(),
		"format":             carapace.ActionValues("json", "pretty"),
		"kubernetes-ca-cert": action.ActionOptionalFiles(),
		"token-locality":     carapace.ActionValues("local", "global"),
	})
}
