package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var create_secret_tlsCmd = &cobra.Command{
	Use:   "tls",
	Short: "Create a TLS secret",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_secret_tlsCmd).Standalone()

	create_secret_tlsCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	create_secret_tlsCmd.Flags().Bool("append-hash", false, "Append a hash of the secret to its name.")
	create_secret_tlsCmd.Flags().String("cert", "", "Path to PEM encoded public key certificate.")
	create_secret_tlsCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	create_secret_tlsCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	create_secret_tlsCmd.Flags().String("generator", "", "The name of the API generator to use.")
	create_secret_tlsCmd.Flags().String("key", "", "Path to private key associated with given certificate.")
	create_secret_tlsCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	create_secret_tlsCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotat")
	create_secret_tlsCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	create_secret_tlsCmd.Flags().Bool("validate", false, "If true, use a schema to validate the input before sending it")
	create_secretCmd.AddCommand(create_secret_tlsCmd)

	carapace.Gen(create_secret_tlsCmd).FlagCompletion(carapace.ActionMap{
		"cert":     carapace.ActionFiles(),
		"dry-run":  action.ActionDryRunModes(),
		"key":      carapace.ActionFiles(),
		"output":   action.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
	})
}
