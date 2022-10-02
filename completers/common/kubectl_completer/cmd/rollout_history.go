package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rollout_historyCmd = &cobra.Command{
	Use:   "history",
	Short: "View rollout history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rollout_historyCmd).Standalone()

	rollout_historyCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only app")
	rollout_historyCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files identifying the resource to get from a server.")
	rollout_historyCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	rollout_historyCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|js")
	rollout_historyCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	rollout_historyCmd.Flags().String("revision", "", "See the details, including podTemplate of the revision specified")
	rollout_historyCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The templa")
	rolloutCmd.AddCommand(rollout_historyCmd)

	carapace.Gen(rollout_historyCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(rollout_historyCmd).PositionalCompletion(
		action.ActionApiResourceResources(),
	)
}
