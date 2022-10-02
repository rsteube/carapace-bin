package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rollout_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the status of the rollout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rollout_statusCmd).Standalone()

	rollout_statusCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files identifying the resource to get from a server.")
	rollout_statusCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	rollout_statusCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related man")
	rollout_statusCmd.Flags().String("revision", "", "Pin to a specific revision for showing its status. Defaults to 0 (last revision).")
	rollout_statusCmd.Flags().String("timeout", "", "The length of time to wait before ending watch, zero means never. Any other values should contain a ")
	rollout_statusCmd.Flags().BoolP("watch", "w", false, "Watch the status of the rollout until it's done.")
	rolloutCmd.AddCommand(rollout_statusCmd)

	carapace.Gen(rollout_statusCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
	})

	carapace.Gen(rollout_statusCmd).PositionalCompletion(
		action.ActionApiResourceResources(),
	)
}
