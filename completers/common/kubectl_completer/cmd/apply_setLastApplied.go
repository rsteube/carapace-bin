package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var setLastAppliedCmd = &cobra.Command{
	Use:   "set-last-applied",
	Short: "Set the last-applied-configuration annotation on a live object to match the contents of a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setLastAppliedCmd).Standalone()

	setLastAppliedCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	setLastAppliedCmd.Flags().Bool("create-annotation", false, "Will create 'last-applied-configuration' annotations if current objects doesn't have one")
	setLastAppliedCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	setLastAppliedCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files that contains the last-applied-configuration annotations")
	setLastAppliedCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	setLastAppliedCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	applyCmd.AddCommand(setLastAppliedCmd)

	carapace.Gen(setLastAppliedCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  action.ActionDryRunModes(),
		"filename": carapace.ActionFiles(),
		"output":   action.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
	})
}
