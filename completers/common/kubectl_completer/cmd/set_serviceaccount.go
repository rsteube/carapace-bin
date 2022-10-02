package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var set_serviceaccountCmd = &cobra.Command{
	Use:   "serviceaccount",
	Short: "Update ServiceAccount of a resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_serviceaccountCmd).Standalone()

	set_serviceaccountCmd.Flags().Bool("all", false, "Select all resources, including uninitialized ones, in the namespace of the specified resource types")
	set_serviceaccountCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	set_serviceaccountCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	set_serviceaccountCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	set_serviceaccountCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files identifying the resource to get from a server.")
	set_serviceaccountCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	set_serviceaccountCmd.Flags().Bool("local", false, "If true, set serviceaccount will NOT contact api-server but run locally.")
	set_serviceaccountCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	set_serviceaccountCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the comman")
	set_serviceaccountCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	set_serviceaccountCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	setCmd.AddCommand(set_serviceaccountCmd)

	carapace.Gen(set_serviceaccountCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   action.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(set_serviceaccountCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if set_serviceaccountCmd.Flag("filename").Changed {
				return action.ActionResources("", "serviceaccounts")
			} else {
				return action.ActionApiResources()
			}
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if set_serviceaccountCmd.Flag("filename").Changed {
				return carapace.ActionValues()
			} else {
				return action.ActionResources("", c.Args[0])
			}
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if set_serviceaccountCmd.Flag("filename").Changed {
				return carapace.ActionValues()
			} else {
				return action.ActionResources("", "serviceaccounts")
			}
		}),
	)
}
