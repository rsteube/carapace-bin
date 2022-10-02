package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var upgrade_diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Show what differences would be applied to existing static pod manifests. See also: kubeadm upgrade apply --dry-run",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_diffCmd).Standalone()
	upgrade_diffCmd.Flags().String("api-server-manifest", "/etc/kubernetes/manifests/kube-apiserver.yaml", "path to API server manifest")
	upgrade_diffCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	upgrade_diffCmd.Flags().IntP("context-lines", "c", 3, "How many lines of context in the diff")
	upgrade_diffCmd.Flags().String("controller-manager-manifest", "/etc/kubernetes/manifests/kube-controller-manager.yaml", "path to controller manifest")
	upgrade_diffCmd.Flags().String("kubeconfig", "/etc/kubernetes/admin.conf", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	upgrade_diffCmd.Flags().String("scheduler-manifest", "/etc/kubernetes/manifests/kube-scheduler.yaml", "path to scheduler manifest")
	upgradeCmd.AddCommand(upgrade_diffCmd)

	carapace.Gen(upgrade_diffCmd).FlagCompletion(carapace.ActionMap{
		"config":                      carapace.ActionFiles(),
		"controller-manager-manifest": carapace.ActionFiles(),
		"kubeconfig":                  carapace.ActionFiles(),
		"scheduler-manifest":          carapace.ActionFiles(),
	})
}
