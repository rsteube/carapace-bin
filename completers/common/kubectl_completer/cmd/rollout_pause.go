package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rollout_pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Mark the provided resource as paused",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rollout_pauseCmd).Standalone()

	rollout_pauseCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	rollout_pauseCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	rollout_pauseCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files identifying the resource to get from a server.")
	rollout_pauseCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	rollout_pauseCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	rollout_pauseCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	rollout_pauseCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	rolloutCmd.AddCommand(rollout_pauseCmd)

	carapace.Gen(rollout_pauseCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(rollout_pauseCmd).PositionalCompletion(
		action.ActionApiResourceResources(),
	)
}
