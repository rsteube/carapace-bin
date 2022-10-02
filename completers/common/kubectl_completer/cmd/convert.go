package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert config files between different API versions. Both YAML and JSON formats are accepted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(convertCmd).Standalone()

	convertCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template..")
	convertCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files to need to get converted.")
	convertCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	convertCmd.Flags().Bool("local", false, "If true, convert will NOT try to contact api-server but run locally.")
	convertCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-as-json|jsonpath-file.")
	convertCmd.Flags().String("output-version", "", "Output the formatted object with the given group version (for ex: 'extensions/v1beta1').")
	convertCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	convertCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file.")
	convertCmd.Flags().Bool("validate", false, "If true, use a schema to validate the input before sending it")
	rootCmd.AddCommand(convertCmd)

	carapace.Gen(convertCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})
}
