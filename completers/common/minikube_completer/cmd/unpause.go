package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var unpauseCmd = &cobra.Command{
	Use:   "unpause",
	Short: "unpause Kubernetes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	unpauseCmd.Flags().BoolP("all-namespaces", "A", false, "If set, unpause all namespaces")
	unpauseCmd.Flags().StringSliceP("namespaces", "n", []string{"kube-system", "kubernetes-dashboard", "storage-gluster", "istio-operator"}, "namespaces to unpause")
	unpauseCmd.Flags().StringP("output", "o", "text", "Format to print stdout in. Options include: [text,json]")
	rootCmd.AddCommand(unpauseCmd)

	carapace.Gen(unpauseCmd).FlagCompletion(carapace.ActionMap{
		"namespaces": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionNamespaces().Invoke(c).Filter(c.Parts).ToA()
		}),
		"output": carapace.ActionValues("text", "json"),
	})
}
