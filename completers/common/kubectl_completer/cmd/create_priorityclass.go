package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var create_priorityclassCmd = &cobra.Command{
	Use:   "priorityclass",
	Short: "Create a priorityclass with the specified name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_priorityclassCmd).Standalone()

	create_priorityclassCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	create_priorityclassCmd.Flags().String("description", "", "description is an arbitrary string that usually provides guidelines on when this priority class shou")
	create_priorityclassCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	create_priorityclassCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	create_priorityclassCmd.Flags().Bool("global-default", false, "global-default specifies whether this PriorityClass should be considered as the default priority.")
	create_priorityclassCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	create_priorityclassCmd.Flags().String("preemption-policy", "", "preemption-policy is the policy for preempting pods with lower priority.")
	create_priorityclassCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotat")
	create_priorityclassCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	create_priorityclassCmd.Flags().Bool("validate", false, "If true, use a schema to validate the input before sending it")
	create_priorityclassCmd.Flags().String("value", "", "the value of this priority class.")
	createCmd.AddCommand(create_priorityclassCmd)

	carapace.Gen(create_priorityclassCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  action.ActionDryRunModes(),
		"output":   action.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
	})
}
