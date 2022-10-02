package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var create_secret_genericCmd = &cobra.Command{
	Use:   "generic",
	Short: "Create a secret from a local file, directory or literal value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_secret_genericCmd).Standalone()

	create_secret_genericCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	create_secret_genericCmd.Flags().Bool("append-hash", false, "Append a hash of the secret to its name.")
	create_secret_genericCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	create_secret_genericCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	create_secret_genericCmd.Flags().String("from-env-file", "", "Specify the path to a file to read lines of key=val pairs to create a secret (i.e. a Docker .env fil")
	create_secret_genericCmd.Flags().String("from-file", "", "Key files can be specified using their file path, in which case a default name will be given to them")
	create_secret_genericCmd.Flags().String("from-literal", "", "Specify a key and literal value to insert in secret (i.e. mykey=somevalue)")
	create_secret_genericCmd.Flags().String("generator", "", "The name of the API generator to use.")
	create_secret_genericCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	create_secret_genericCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotat")
	create_secret_genericCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	create_secret_genericCmd.Flags().String("type", "", "The type of secret to create")
	create_secret_genericCmd.Flags().Bool("validate", false, "If true, use a schema to validate the input before sending it")
	create_secretCmd.AddCommand(create_secret_genericCmd)

	carapace.Gen(create_secret_genericCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":       action.ActionDryRunModes(),
		"from-env-file": carapace.ActionFiles(),
		"from-file":     carapace.ActionFiles(),
		"output":        action.ActionOutputFormats(),
		"template":      carapace.ActionFiles(),
	})
}
