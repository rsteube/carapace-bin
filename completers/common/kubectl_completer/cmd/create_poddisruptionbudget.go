package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var create_poddisruptionbudgetCmd = &cobra.Command{
	Use:   "poddisruptionbudget",
	Short: "Create a pod disruption budget with the specified name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_poddisruptionbudgetCmd).Standalone()

	create_poddisruptionbudgetCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	create_poddisruptionbudgetCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	create_poddisruptionbudgetCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	create_poddisruptionbudgetCmd.Flags().String("generator", "", "The name of the API generator to use.")
	create_poddisruptionbudgetCmd.Flags().String("max-unavailable", "", "The maximum number or percentage of unavailable pods this budget requires.")
	create_poddisruptionbudgetCmd.Flags().String("min-available", "", "The minimum number or percentage of available pods this budget requires.")
	create_poddisruptionbudgetCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	create_poddisruptionbudgetCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotat")
	create_poddisruptionbudgetCmd.Flags().String("selector", "", "A label selector to use for this budget. Only equality-based selector requirements are supported.")
	create_poddisruptionbudgetCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	create_poddisruptionbudgetCmd.Flags().Bool("validate", false, "If true, use a schema to validate the input before sending it")
	createCmd.AddCommand(create_poddisruptionbudgetCmd)

	carapace.Gen(create_poddisruptionbudgetCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  action.ActionDryRunModes(),
		"output":   action.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
	})
}
