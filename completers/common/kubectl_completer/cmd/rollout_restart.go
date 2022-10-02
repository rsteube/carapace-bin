package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rollout_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart a resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rollout_restartCmd).Standalone()

	rollout_restartCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	rollout_restartCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	rollout_restartCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files identifying the resource to get from a server.")
	rollout_restartCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	rollout_restartCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	rollout_restartCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	rollout_restartCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	rolloutCmd.AddCommand(rollout_restartCmd)

	carapace.Gen(rollout_restartCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(rollout_restartCmd).PositionalCompletion(
		action.ActionApiResourceResources(),
	)
}
