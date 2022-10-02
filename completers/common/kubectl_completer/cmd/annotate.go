package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var annotateCmd = &cobra.Command{
	Use:   "annotate",
	Short: "Update the annotations on a resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(annotateCmd).Standalone()

	annotateCmd.Flags().Bool("all", false, "Select all resources, including uninitialized ones, in the namespace of the specified resource types")
	annotateCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	annotateCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	annotateCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	annotateCmd.Flags().String("field-selector", "", "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1")
	annotateCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files identifying the resource to update the annotation")
	annotateCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	annotateCmd.Flags().Bool("list", false, "If true, display the annotations for a given resource.")
	annotateCmd.Flags().Bool("local", false, "If true, annotation will NOT contact api-server but run locally.")
	annotateCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	annotateCmd.Flags().Bool("overwrite", false, "If true, allow annotations to be overwritten, otherwise reject annotation updates that overwrite exi")
	annotateCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the comman")
	annotateCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	annotateCmd.Flags().String("resource-version", "", "If non-empty, the annotation update will only succeed if this is the current resource-version for th")
	annotateCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, not including uninitialized ones, supports '=', '==', and '!='.")
	annotateCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	rootCmd.AddCommand(annotateCmd)

	carapace.Gen(annotateCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   action.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(annotateCmd).PositionalCompletion(
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

	carapace.Gen(annotateCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if labelCmd.Flag("filename").Changed || labelCmd.Flag("all").Changed {
				return carapace.ActionValues() // TODO support labels for file
			} else {
				return action.ActionAnnotations("", c.Args[0]+"/"+c.Args[1])
			}
		}),
	)
}
