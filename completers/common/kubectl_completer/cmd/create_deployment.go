package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var create_deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Create a cronjob with the specified name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_deploymentCmd).Standalone()

	create_deploymentCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	create_deploymentCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	create_deploymentCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	create_deploymentCmd.Flags().String("generator", "", "The name of the API generator to use.")
	create_deploymentCmd.Flags().String("image", "", "Image names to run.")
	create_deploymentCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	create_deploymentCmd.Flags().String("port", "", "The port that this container exposes.")
	create_deploymentCmd.Flags().StringP("replicas", "r", "", "Number of replicas to create. Default is 1.")
	create_deploymentCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotat")
	create_deploymentCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	create_deploymentCmd.Flags().Bool("validate", false, "If true, use a schema to validate the input before sending it")
	createCmd.AddCommand(create_deploymentCmd)

	carapace.Gen(create_deploymentCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  action.ActionDryRunModes(),
		"output":   action.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
	})
}
