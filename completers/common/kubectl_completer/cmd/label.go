package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var labelCmd = &cobra.Command{
	Use:   "label",
	Short: "Update the labels on a resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(labelCmd).Standalone()

	labelCmd.Flags().Bool("all", false, "Select all resources, including uninitialized ones, in the namespace of the specified resource types")
	labelCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	labelCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	labelCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	labelCmd.Flags().String("field-selector", "", "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1")
	labelCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files identifying the resource to update the labels")
	labelCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	labelCmd.Flags().Bool("list", false, "If true, display the labels for a given resource.")
	labelCmd.Flags().Bool("local", false, "If true, label will NOT contact api-server but run locally.")
	labelCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	labelCmd.Flags().Bool("overwrite", false, "If true, allow labels to be overwritten, otherwise reject label updates that overwrite existing labe")
	labelCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the comman")
	labelCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	labelCmd.Flags().String("resource-version", "", "If non-empty, the labels update will only succeed if this is the current resource-version for the ob")
	labelCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, not including uninitialized ones, supports '=', '==', and '!='.")
	labelCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	rootCmd.AddCommand(labelCmd)

	carapace.Gen(labelCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   action.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(labelCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if labelCmd.Flag("filename").Changed {
				return carapace.ActionValues() // TODO get labels for file
			} else {
				return action.ActionApiResources()
			}
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if labelCmd.Flag("all").Changed {
				return carapace.ActionValues()
			} else {
				return action.ActionResources("", c.Args[0])
			}
		}),
	)

	carapace.Gen(labelCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if labelCmd.Flag("filename").Changed || labelCmd.Flag("all").Changed {
				return carapace.ActionValues() // TODO support labels for file
			} else {
				return action.ActionLabels("", c.Args[0]+"/"+c.Args[1])
			}
		}),
	)
}
