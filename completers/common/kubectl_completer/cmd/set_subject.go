package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var set_subjectCmd = &cobra.Command{
	Use:   "subject",
	Short: "Update User, Group or ServiceAccount in a RoleBinding/ClusterRoleBinding",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_subjectCmd).Standalone()

	set_subjectCmd.Flags().Bool("all", false, "Select all resources, including uninitialized ones, in the namespace of the specified resource types")
	set_subjectCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	set_subjectCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	set_subjectCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	set_subjectCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files the resource to update the subjects")
	set_subjectCmd.Flags().String("group", "", "Groups to bind to the role")
	set_subjectCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	set_subjectCmd.Flags().Bool("local", false, "If true, set subject will NOT contact api-server but run locally.")
	set_subjectCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	set_subjectCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	set_subjectCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, not including uninitialized ones, supports '=', '==', and '!='.")
	set_subjectCmd.Flags().String("serviceaccount", "", "Service accounts to bind to the role")
	set_subjectCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	setCmd.AddCommand(set_subjectCmd)

	carapace.Gen(set_subjectCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":        action.ActionDryRunModes(),
		"filename":       carapace.ActionFiles(),
		"kustomize":      carapace.ActionDirectories(),
		"output":         action.ActionOutputFormats(),
		"serviceaccount": action.ActionNamespaceServiceAccounts(),
		"template":       carapace.ActionFiles(),
	})

	carapace.Gen(set_subjectCmd).PositionalCompletion(
		carapace.ActionValues("rolebinding", "clusterrolebinding"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionResources("", c.Args[0])
		}),
	)
}
