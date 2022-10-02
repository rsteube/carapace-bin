package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a resource on the server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(editCmd).Standalone()

	editCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	editCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	editCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files to use to edit the resource")
	editCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	editCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	editCmd.Flags().Bool("output-patch", false, "Output the patch if the resource is edited.")
	editCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the comman")
	editCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	editCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotat")
	editCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	editCmd.Flags().Bool("validate", false, "If true, use a schema to validate the input before sending it")
	editCmd.Flags().Bool("windows-line-endings", false, "Defaults to the line ending native to your platform.")
	rootCmd.AddCommand(editCmd)

	carapace.Gen(editCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(editCmd).PositionalCompletion(
		action.ActionApiResourceResources(),
	)
}
