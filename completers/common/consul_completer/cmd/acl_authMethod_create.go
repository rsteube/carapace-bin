package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var acl_authMethod_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an ACL auth method",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_authMethod_createCmd).Standalone()
	addClientFlags(acl_authMethod_createCmd)
	addServerFlags(acl_authMethod_createCmd)

	acl_authMethod_createCmd.Flags().String("config", "", "The configuration for the auth method.")
	acl_authMethod_createCmd.Flags().String("description", "", "A description of the auth method.")
	acl_authMethod_createCmd.Flags().String("display-name", "", "An optional name to use instead of the name when displaying this auth method in a UI.")
	acl_authMethod_createCmd.Flags().String("format", "", "Output format.")
	acl_authMethod_createCmd.Flags().String("kubernetes-ca-cert", "", "PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.")
	acl_authMethod_createCmd.Flags().String("kubernetes-host", "", "Address of the Kubernetes API server.")
	acl_authMethod_createCmd.Flags().String("kubernetes-service-account-jwt", "", "A Kubernetes service account JWT used to access the TokenReview API to validate other JWTs during login.")
	acl_authMethod_createCmd.Flags().String("max-token-ttl", "", "Duration of time all tokens created by this auth method should be valid for")
	acl_authMethod_createCmd.Flags().Bool("meta", false, "Indicates that auth method metadata such as the raft indices should be shown for each entry.")
	acl_authMethod_createCmd.Flags().String("name", "", "The new auth method's name.")
	acl_authMethod_createCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_authMethod_createCmd.Flags().String("token-locality", "", "Defines the kind of token that this auth method should produce.")
	acl_authMethod_createCmd.Flags().String("type", "", "The new auth method's type.")
	acl_authMethodCmd.AddCommand(acl_authMethod_createCmd)

	carapace.Gen(acl_authMethod_createCmd).FlagCompletion(carapace.ActionMap{
		// TODO flag completions
		"config":             action.ActionOptionalFiles(),
		"format":             carapace.ActionValues("json", "pretty"),
		"kubernetes-ca-cert": action.ActionOptionalFiles(),
	})
}
