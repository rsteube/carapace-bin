package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var create_clusterrolebindingCmd = &cobra.Command{
	Use:   "clusterrolebinding",
	Short: "Create a ClusterRoleBinding for a particular ClusterRole",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_clusterrolebindingCmd).Standalone()

	create_clusterrolebindingCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	create_clusterrolebindingCmd.Flags().String("clusterrole", "", "ClusterRole this ClusterRoleBinding should reference")
	create_clusterrolebindingCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	create_clusterrolebindingCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	create_clusterrolebindingCmd.Flags().String("group", "", "Groups to bind to the clusterrole")
	create_clusterrolebindingCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	create_clusterrolebindingCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotat")
	create_clusterrolebindingCmd.Flags().String("serviceaccount", "", "Service accounts to bind to the clusterrole, in the format <namespace>:<name>")
	create_clusterrolebindingCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	create_clusterrolebindingCmd.Flags().Bool("validate", false, "If true, use a schema to validate the input before sending it")
	createCmd.AddCommand(create_clusterrolebindingCmd)

	carapace.Gen(create_clusterrolebindingCmd).FlagCompletion(carapace.ActionMap{
		"clusterrole":    action.ActionResources("", "clusterrole"),
		"dry-run":        action.ActionDryRunModes(),
		"output":         action.ActionOutputFormats(),
		"serviceaccount": action.ActionNamespaceServiceAccounts(),
		"template":       carapace.ActionFiles(),
	})
}
