package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var auth_reconsileCmd = &cobra.Command{
	Use:   "reconsile",
	Short: "Reconciles rules for RBAC Role, RoleBinding, ClusterRole, and ClusterRole binding objects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_reconsileCmd).Standalone()

	auth_reconsileCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	auth_reconsileCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	auth_reconsileCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files identifying the resource to reconcile.")
	auth_reconsileCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	auth_reconsileCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	auth_reconsileCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	auth_reconsileCmd.Flags().Bool("remove-extra-permissions", false, "If true, removes extra permissions added to roles")
	auth_reconsileCmd.Flags().Bool("remove-extra-subjects", false, "If true, removes extra subjects added to rolebindings")
	auth_reconsileCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	authCmd.AddCommand(auth_reconsileCmd)

	carapace.Gen(auth_reconsileCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   action.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})
}
