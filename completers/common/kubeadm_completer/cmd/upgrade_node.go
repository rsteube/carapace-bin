package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubeadm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var upgrade_nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Upgrade commands for a node in the cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_nodeCmd).Standalone()
	upgrade_nodeCmd.Flags().Bool("certificate-renewal", true, "Perform the renewal of certificates used by component changed during upgrades.")
	upgrade_nodeCmd.Flags().Bool("dry-run", false, "Do not change any state, just output the actions that would be performed.")
	upgrade_nodeCmd.Flags().Bool("etcd-upgrade", true, "Perform the upgrade of etcd.")
	upgrade_nodeCmd.Flags().StringSlice("ignore-preflight-errors", []string{}, "A list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.")
	upgrade_nodeCmd.Flags().String("kubeconfig", "/etc/kubernetes/kubelet.conf", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	upgrade_nodeCmd.Flags().String("patches", "", "Path to a directory that contains files named \"target[suffix][+patchtype].extension\".")
	upgrade_nodeCmd.Flags().StringSlice("skip-phases", []string{}, "List of phases to be skipped")
	upgradeCmd.AddCommand(upgrade_nodeCmd)

	carapace.Gen(upgrade_nodeCmd).FlagCompletion(carapace.ActionMap{
		"ignore-preflight-errors": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionChecks().Invoke(c).Filter(c.Parts).ToA()
		}),
		"kubeconfig": carapace.ActionFiles(),
		"patches":    carapace.ActionDirectories(),
		"skip-phases": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionPhases().Invoke(c).Filter(c.Parts).ToMultiPartsA("/")
		}),
	})
}
