package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var create_cronjobCmd = &cobra.Command{
	Use:   "cronjob",
	Short: "Create a cronjob with the specified name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_cronjobCmd).Standalone()

	create_cronjobCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	create_cronjobCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	create_cronjobCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	create_cronjobCmd.Flags().String("image", "", "Image name to run.")
	create_cronjobCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	create_cronjobCmd.Flags().String("restart", "", "job's restart policy. supported values: OnFailure, Never")
	create_cronjobCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotat")
	create_cronjobCmd.Flags().String("schedule", "", "A schedule in the Cron format the job should be run with.")
	create_cronjobCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	create_cronjobCmd.Flags().Bool("validate", false, "If true, use a schema to validate the input before sending it")
	createCmd.AddCommand(create_cronjobCmd)

	carapace.Gen(create_cronjobCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  action.ActionDryRunModes(),
		"output":   action.ActionOutputFormats(),
		"restart":  carapace.ActionValues("OnFailure", "Never"),
		"template": carapace.ActionFiles(),
	})
}
