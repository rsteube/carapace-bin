package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var set_envCmd = &cobra.Command{
	Use:   "env",
	Short: "Update environment variables on a pod template",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_envCmd).Standalone()

	set_envCmd.Flags().Bool("all", false, "If true, select all resources in the namespace of the specified resource types")
	set_envCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	set_envCmd.Flags().StringP("containers", "c", "", "The names of containers in the selected pod templates to change - may use wildcards")
	set_envCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	set_envCmd.Flags().StringP("env", "e", "", "Specify a key-value pair for an environment variable to set into each container.")
	set_envCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	set_envCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files the resource to update the env")
	set_envCmd.Flags().String("from", "", "The name of a resource from which to inject environment variables")
	set_envCmd.Flags().String("keys", "", "Comma-separated list of keys to import from specified resource")
	set_envCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	set_envCmd.Flags().Bool("list", false, "If true, display the environment and any changes in the standard format. this flag will removed when")
	set_envCmd.Flags().Bool("local", false, "If true, set env will NOT contact api-server but run locally.")
	set_envCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	set_envCmd.Flags().Bool("overwrite", false, "If true, allow environment to be overwritten, otherwise reject updates that overwrite existing envir")
	set_envCmd.Flags().String("prefix", "", "Prefix to append to variable names")
	set_envCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	set_envCmd.Flags().Bool("resolve", false, "If true, show secret or configmap references when listing variables")
	set_envCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on")
	set_envCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	setCmd.AddCommand(set_envCmd)

	carapace.Gen(set_envCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  action.ActionDryRunModes(),
		"filename": carapace.ActionFiles(),
		"from":     action.ActionApiResourceResources(),
		"output":   action.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
	})

	carapace.Gen(set_envCmd).PositionalCompletion(
		action.ActionApiResourceResources(),
	)
}
