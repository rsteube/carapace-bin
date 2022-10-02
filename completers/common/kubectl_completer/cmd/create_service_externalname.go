package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var create_service_externalnameCmd = &cobra.Command{
	Use:   "externalname",
	Short: "Create an ExternalName service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_service_externalnameCmd).Standalone()

	create_service_externalnameCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	create_service_externalnameCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	create_service_externalnameCmd.Flags().String("external-name", "", "External name of service")
	create_service_externalnameCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	create_service_externalnameCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	create_service_externalnameCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotat")
	create_service_externalnameCmd.Flags().String("tcp", "", "Port pairs can be specified as '<port>:<targetPort>'.")
	create_service_externalnameCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	create_service_externalnameCmd.Flags().Bool("validate", false, "If true, use a schema to validate the input before sending it")
	create_serviceCmd.AddCommand(create_service_externalnameCmd)

	carapace.Gen(create_service_externalnameCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  action.ActionDryRunModes(),
		"output":   action.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
	})
}
