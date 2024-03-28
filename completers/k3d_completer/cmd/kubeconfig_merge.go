package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubeconfig_mergeCmd = &cobra.Command{
	Use:     "merge [CLUSTER [CLUSTER [...]] | --all]",
	Short:   "Write/Merge kubeconfig(s) from cluster(s) into new or existing kubeconfig/file.",
	Aliases: []string{"write"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubeconfig_mergeCmd).Standalone()

	kubeconfig_mergeCmd.Flags().BoolP("all", "a", false, "Get kubeconfigs from all existing clusters")
	kubeconfig_mergeCmd.Flags().BoolP("kubeconfig-merge-default", "d", false, "Merge into the default kubeconfig ($KUBECONFIG or /home/rsteube/.kube/config)")
	kubeconfig_mergeCmd.Flags().BoolP("kubeconfig-switch-context", "s", false, "Switch to new context")
	kubeconfig_mergeCmd.Flags().StringP("output", "o", "", "Define output [ - | FILE ] (default from $KUBECONFIG or /home/rsteube/.kube/config")
	kubeconfig_mergeCmd.Flags().Bool("overwrite", false, "[Careful!] Overwrite existing file, ignoring its contents")
	kubeconfig_mergeCmd.Flags().BoolP("update", "u", false, "Update conflicting fields in existing kubeconfig")
	kubeconfigCmd.AddCommand(kubeconfig_mergeCmd)
}
