package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var taintCmd = &cobra.Command{
	Use:   "taint",
	Short: "Update the taints on one or more nodes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taintCmd).Standalone()

	taintCmd.Flags().Bool("all", false, "Select all nodes in the cluster")
	taintCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	taintCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	taintCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	taintCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	taintCmd.Flags().Bool("overwrite", false, "If true, allow taints to be overwritten, otherwise reject taint updates that overwrite existing tain")
	taintCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)")
	taintCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	taintCmd.Flags().Bool("validate", false, "If true, use a schema to validate the input before sending it")
	rootCmd.AddCommand(taintCmd)

	carapace.Gen(taintCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  action.ActionDryRunModes(),
		"output":   action.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
	})

	carapace.Gen(taintCmd).PositionalCompletion(
		carapace.ActionValues("nodes"),
		action.ActionResources("", "nodes"),
	)
}
