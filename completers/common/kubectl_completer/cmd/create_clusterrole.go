package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var create_clusterroleCmd = &cobra.Command{
	Use:   "clusterrole",
	Short: "Create a ClusterRole",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_clusterroleCmd).Standalone()

	create_clusterroleCmd.Flags().String("aggregation-rule", "", "An aggregation label selector for combining ClusterRoles.")
	create_clusterroleCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	create_clusterroleCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	create_clusterroleCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	create_clusterroleCmd.Flags().String("non-resource-url", "", "A partial url that user should have access to.")
	create_clusterroleCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	create_clusterroleCmd.Flags().String("resource", "", "Resource that the rule applies to")
	create_clusterroleCmd.Flags().String("resource-name", "", "Resource in the white list that the rule applies to, repeat this flag for multiple items")
	create_clusterroleCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotat")
	create_clusterroleCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	create_clusterroleCmd.Flags().Bool("validate", false, "If true, use a schema to validate the input before sending it")
	create_clusterroleCmd.Flags().String("verb", "", "Verb that applies to the resources contained in the rule")
	createCmd.AddCommand(create_clusterroleCmd)

	carapace.Gen(create_clusterroleCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  action.ActionDryRunModes(),
		"output":   action.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
		"verb":     action.ActionResourceVerbs(),
	})
}
