package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var create_jobCmd = &cobra.Command{
	Use:   "job",
	Short: "Create a job with the specified name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_jobCmd).Standalone()

	create_jobCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	create_jobCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	create_jobCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	create_jobCmd.Flags().String("from", "", "The name of the resource to create a Job from (only cronjob is supported).")
	create_jobCmd.Flags().String("image", "", "Image name to run.")
	create_jobCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	create_jobCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotat")
	create_jobCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	create_jobCmd.Flags().Bool("validate", false, "If true, use a schema to validate the input before sending it")
	createCmd.AddCommand(create_jobCmd)

	carapace.Gen(create_jobCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  action.ActionDryRunModes(),
		"from":     action.ActionResources("", "cronjobs"),
		"output":   action.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
	})
}
