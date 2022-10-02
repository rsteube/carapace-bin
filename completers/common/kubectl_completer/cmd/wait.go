package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var waitCmd = &cobra.Command{
	Use:   "wait",
	Short: "Experimental: Wait for a specific condition on one or many resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitCmd).Standalone()

	waitCmd.Flags().Bool("all", false, "Select all resources in the namespace of the specified resource types")
	waitCmd.Flags().BoolP("all-namespaces", "A", false, "If present, list the requested object(s) across all namespaces. Namespace in current context is igno")
	waitCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	waitCmd.Flags().String("field-selector", "", "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1")
	waitCmd.Flags().StringP("filename", "f", "", "identifying the resource.")
	waitCmd.Flags().String("for", "", "The condition to wait on: [delete|condition=condition-name].")
	waitCmd.Flags().Bool("local", false, "If true, annotation will NOT contact api-server but run locally.")
	waitCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	waitCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	waitCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)")
	waitCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	waitCmd.Flags().String("timeout", "", "The length of time to wait before giving up.  Zero means check once and don't wait, negative means w")
	rootCmd.AddCommand(waitCmd)

	carapace.Gen(waitCmd).FlagCompletion(carapace.ActionMap{
		"filename": carapace.ActionFiles(),
		"output":   action.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
	})
}
