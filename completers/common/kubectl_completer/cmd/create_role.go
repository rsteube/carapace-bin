package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var create_roleCmd = &cobra.Command{
	Use:   "role",
	Short: "Create a role with single rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_roleCmd).Standalone()

	create_roleCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	create_roleCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	create_roleCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	create_roleCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	create_roleCmd.Flags().String("resource", "", "Resource that the rule applies to")
	create_roleCmd.Flags().String("resource-name", "", "Resource in the white list that the rule applies to, repeat this flag for multiple items")
	create_roleCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotat")
	create_roleCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	create_roleCmd.Flags().Bool("validate", false, "If true, use a schema to validate the input before sending it")
	create_roleCmd.Flags().String("verb", "", "Verb that applies to the resources contained in the rule")
	createCmd.AddCommand(create_roleCmd)

	carapace.Gen(create_roleCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  action.ActionDryRunModes(),
		"output":   action.ActionOutputFormats(),
		"resource": action.ActionApiResources(),
		"template": carapace.ActionFiles(),
		"verb":     action.ActionResourceVerbs(),
	})
}
