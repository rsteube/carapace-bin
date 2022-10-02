package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var apply_editLastAppliedCmd = &cobra.Command{
	Use:   "edit-last-applied",
	Short: "Edit latest last-applied-configuration annotations of a resource/object",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apply_editLastAppliedCmd).Standalone()

	apply_editLastAppliedCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	apply_editLastAppliedCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	apply_editLastAppliedCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files to use to edit the resource")
	apply_editLastAppliedCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	apply_editLastAppliedCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	apply_editLastAppliedCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the comman")
	apply_editLastAppliedCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	apply_editLastAppliedCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	apply_editLastAppliedCmd.Flags().Bool("windows-line-endings", false, "Defaults to the line ending native to your platform.")
	applyCmd.AddCommand(apply_editLastAppliedCmd)

	carapace.Gen(apply_editLastAppliedCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(apply_editLastAppliedCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if apply_editLastAppliedCmd.Flag("filename").Changed {
				return carapace.ActionValues()
			} else {
				return action.ActionApiResourceResources()
			}
		}),
	)
}
