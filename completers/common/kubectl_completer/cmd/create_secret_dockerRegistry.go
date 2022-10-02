package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var create_secret_dockerRegistryCmd = &cobra.Command{
	Use:   "docker-registry",
	Short: "Create a secret for use with a Docker registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_secret_dockerRegistryCmd).Standalone()

	create_secret_dockerRegistryCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	create_secret_dockerRegistryCmd.Flags().Bool("append-hash", false, "Append a hash of the secret to its name.")
	create_secret_dockerRegistryCmd.Flags().String("docker-email", "", "Email for Docker registry")
	create_secret_dockerRegistryCmd.Flags().String("docker-password", "", "Password for Docker registry authentication")
	create_secret_dockerRegistryCmd.Flags().String("docker-server", "", "Server location for Docker registry")
	create_secret_dockerRegistryCmd.Flags().String("docker-username", "", "Username for Docker registry authentication")
	create_secret_dockerRegistryCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	create_secret_dockerRegistryCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	create_secret_dockerRegistryCmd.Flags().String("from-file", "", "Key files can be specified using their file path, in which case a default name will be given to them")
	create_secret_dockerRegistryCmd.Flags().String("generator", "", "The name of the API generator to use.")
	create_secret_dockerRegistryCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	create_secret_dockerRegistryCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotat")
	create_secret_dockerRegistryCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	create_secret_dockerRegistryCmd.Flags().Bool("validate", false, "If true, use a schema to validate the input before sending it")
	create_secretCmd.AddCommand(create_secret_dockerRegistryCmd)

	carapace.Gen(create_secret_dockerRegistryCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   action.ActionDryRunModes(),
		"from-file": carapace.ActionFiles(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})
}
