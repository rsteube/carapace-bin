package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Update field(s) of a resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(patchCmd).Standalone()

	patchCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	patchCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	patchCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	patchCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files identifying the resource to update")
	patchCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	patchCmd.Flags().Bool("local", false, "If true, patch will operate on the content of the file, not the server-side resource.")
	patchCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	patchCmd.Flags().StringP("patch", "p", "", "The patch to be applied to the resource JSON file.")
	patchCmd.Flags().String("patch-file", "", "A file containing a patch to be applied to the resource.")
	patchCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the comman")
	patchCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	patchCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	patchCmd.Flags().String("type", "", "The type of patch being provided; one of [json merge strategic]")
	rootCmd.AddCommand(patchCmd)

	carapace.Gen(patchCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":    action.ActionDryRunModes(),
		"filename":   carapace.ActionFiles(),
		"kustomize":  carapace.ActionDirectories(),
		"output":     action.ActionOutputFormats(),
		"patch-file": carapace.ActionFiles(),
		"template":   carapace.ActionFiles(),
		"type":       carapace.ActionValues("json", "merge", "strategic"),
	})

	carapace.Gen(patchCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if patchCmd.Flag("filename").Changed {
				return carapace.ActionValues()
			} else {
				return action.ActionApiResources()
			}
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if patchCmd.Flag("filename").Changed {
				return carapace.ActionValues()
			} else {
				return action.ActionResources("", c.Args[0])
			}
		}),
	)
}
