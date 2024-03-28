package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubeconfig_getCmd = &cobra.Command{
	Use:     "get [CLUSTER [CLUSTER [...]] | --all]",
	Short:   "Print kubeconfig(s) from cluster(s).",
	Aliases: []string{"print", "show"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubeconfig_getCmd).Standalone()

	kubeconfig_getCmd.Flags().BoolP("all", "a", false, "Output kubeconfigs from all existing clusters")
	kubeconfigCmd.AddCommand(kubeconfig_getCmd)
}
