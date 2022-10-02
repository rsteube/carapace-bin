package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Display merged kubeconfig settings or a specified kubeconfig file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_viewCmd).Standalone()

	config_viewCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	config_viewCmd.Flags().Bool("flatten", false, "Flatten the resulting kubeconfig file into self-contained output (useful for creating portable kubec")
	config_viewCmd.Flags().String("merge", "", "Merge the full hierarchy of kubeconfig files")
	config_viewCmd.Flags().Bool("minify", false, "Remove all information not used by current-context from the output")
	config_viewCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	config_viewCmd.Flags().Bool("raw", false, "Display raw byte data")
	config_viewCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	configCmd.AddCommand(config_viewCmd)

	carapace.Gen(config_viewCmd).FlagCompletion(carapace.ActionMap{
		"output":   action.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
	})
}
