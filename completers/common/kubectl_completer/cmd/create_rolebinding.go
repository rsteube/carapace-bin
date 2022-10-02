package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var create_rolebindingCmd = &cobra.Command{
	Use:   "rolebinding",
	Short: "Create a RoleBinding for a particular Role or ClusterRole",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_rolebindingCmd).Standalone()

	create_rolebindingCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	create_rolebindingCmd.Flags().String("clusterrole", "", "ClusterRole this RoleBinding should reference")
	create_rolebindingCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	create_rolebindingCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	create_rolebindingCmd.Flags().String("group", "", "Groups to bind to the role")
	create_rolebindingCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	create_rolebindingCmd.Flags().String("role", "", "Role this RoleBinding should reference")
	create_rolebindingCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotat")
	create_rolebindingCmd.Flags().String("serviceaccount", "", "Service accounts to bind to the role, in the format <namespace>:<name>")
	create_rolebindingCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	create_rolebindingCmd.Flags().Bool("validate", false, "If true, use a schema to validate the input before sending it")
	createCmd.AddCommand(create_rolebindingCmd)

	carapace.Gen(create_rolebindingCmd).FlagCompletion(carapace.ActionMap{
		"clusterrole":    action.ActionResources("", "clusterroles"),
		"dry-run":        action.ActionDryRunModes(),
		"output":         action.ActionOutputFormats(),
		"role":           action.ActionResources("", "roles"),
		"serviceaccount": action.ActionNamespaceServiceAccounts(),
		"template":       carapace.ActionFiles(),
	})
}
