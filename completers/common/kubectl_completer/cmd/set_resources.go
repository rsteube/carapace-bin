package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var set_resourcesCmd = &cobra.Command{
	Use:   "resources",
	Short: "Update resource requests/limits on objects with pod templates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_resourcesCmd).Standalone()

	set_resourcesCmd.Flags().Bool("all", false, "Select all resources, including uninitialized ones, in the namespace of the specified resource types")
	set_resourcesCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	set_resourcesCmd.Flags().StringP("containers", "c", "", "The names of containers in the selected pod templates to change, all containers are selected by defa")
	set_resourcesCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	set_resourcesCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	set_resourcesCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files identifying the resource to get from a server.")
	set_resourcesCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	set_resourcesCmd.Flags().String("limits", "", "The resource requirement requests for this container.  For example, 'cpu=100m,memory=256Mi'.  Note t")
	set_resourcesCmd.Flags().Bool("local", false, "If true, set resources will NOT contact api-server but run locally.")
	set_resourcesCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	set_resourcesCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the comman")
	set_resourcesCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	set_resourcesCmd.Flags().String("requests", "", "The resource requirement requests for this container.  For example, 'cpu=100m,memory=256Mi'.  Note t")
	set_resourcesCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, not including uninitialized ones,supports '=', '==', and '!='.(")
	set_resourcesCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	setCmd.AddCommand(set_resourcesCmd)

	carapace.Gen(set_resourcesCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   action.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(set_resourcesCmd).PositionalCompletion(
		carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if set_resourcesCmd.Flag("filename").Changed {
				return carapace.ActionValues()
			} else {
				return action.ActionApiResources()
			}
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if set_resourcesCmd.Flag("filename").Changed {
				return carapace.ActionValues()
			} else {
				return action.ActionResources("", c.Args[0])
			}
		}),
	)
}
