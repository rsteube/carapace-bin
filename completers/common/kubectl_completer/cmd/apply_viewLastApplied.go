package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var apply_viewLastAppliedCmd = &cobra.Command{
	Use:   "view-last-applied",
	Short: "View latest last-applied-configuration annotations of a resource/object",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apply_viewLastAppliedCmd).Standalone()

	apply_viewLastAppliedCmd.Flags().Bool("all", false, "Select all resources in the namespace of the specified resource types")
	apply_viewLastAppliedCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files that contains the last-applied-configuration annotations")
	apply_viewLastAppliedCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	apply_viewLastAppliedCmd.Flags().StringP("output", "o", "", "Output format. Must be one of yaml|json")
	apply_viewLastAppliedCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	apply_viewLastAppliedCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)")
	applyCmd.AddCommand(apply_viewLastAppliedCmd)

	carapace.Gen(apply_viewLastAppliedCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
	})

	carapace.Gen(apply_viewLastAppliedCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if apply_editLastAppliedCmd.Flag("filename").Changed {
				return carapace.ActionValues()
			} else {
				return action.ActionApiResourceResources()
			}
		}),
	)
}
