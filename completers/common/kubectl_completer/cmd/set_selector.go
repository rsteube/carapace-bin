package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var set_selectorCmd = &cobra.Command{
	Use:   "selector",
	Short: "Set the selector on a resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_selectorCmd).Standalone()

	set_selectorCmd.Flags().Bool("all", false, "Select all resources in the namespace of the specified resource types")
	set_selectorCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	set_selectorCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	set_selectorCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	set_selectorCmd.Flags().StringP("filename", "f", "", "identifying the resource.")
	set_selectorCmd.Flags().Bool("local", false, "If true, annotation will NOT contact api-server but run locally.")
	set_selectorCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	set_selectorCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the comman")
	set_selectorCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	set_selectorCmd.Flags().String("resource-version", "", "If non-empty, the selectors update will only succeed if this is the current resource-version for the")
	set_selectorCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	setCmd.AddCommand(set_selectorCmd)

	carapace.Gen(set_selectorCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  action.ActionDryRunModes(),
		"filename": carapace.ActionFiles(),
		"output":   action.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
	})
}
