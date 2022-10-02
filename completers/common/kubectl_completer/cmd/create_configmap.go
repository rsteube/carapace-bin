package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var create_configmapCmd = &cobra.Command{
	Use:   "configmap",
	Short: "Create a configmap from a local file, directory or literal value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_configmapCmd).Standalone()

	create_configmapCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	create_configmapCmd.Flags().Bool("append-hash", false, "Append a hash of the configmap to its name.")
	create_configmapCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	create_configmapCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	create_configmapCmd.Flags().String("from-env-file", "", "Specify the path to a file to read lines of key=val pairs to create a configmap (i.e. a Docker .env ")
	create_configmapCmd.Flags().String("from-file", "", "Key file can be specified using its file path, in which case file basename will be used as configmap")
	create_configmapCmd.Flags().String("from-literal", "", "Specify a key and literal value to insert in configmap (i.e. mykey=somevalue)")
	create_configmapCmd.Flags().String("generator", "", "The name of the API generator to use.")
	create_configmapCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	create_configmapCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotat")
	create_configmapCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	create_configmapCmd.Flags().Bool("validate", false, "If true, use a schema to validate the input before sending it")
	createCmd.AddCommand(create_configmapCmd)

	carapace.Gen(create_configmapCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":       action.ActionDryRunModes(),
		"from-env-file": carapace.ActionFiles(),
		"output":        action.ActionOutputFormats(),
		"template":      carapace.ActionFiles(),
	})
}
