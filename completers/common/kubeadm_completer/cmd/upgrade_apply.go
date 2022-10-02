package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubeadm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var upgrade_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Upgrade your Kubernetes cluster to the specified version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_applyCmd).Standalone()
	upgrade_applyCmd.Flags().Bool("allow-experimental-upgrades", false, "Show unstable versions of Kubernetes as an upgrade alternative and allow upgrading to an alpha/beta/release candidate versions of Kubernetes.")
	upgrade_applyCmd.Flags().Bool("allow-release-candidate-upgrades", false, "Show release candidate versions of Kubernetes as an upgrade alternative and allow upgrading to a release candidate versions of Kubernetes.")
	upgrade_applyCmd.Flags().Bool("certificate-renewal", true, "Perform the renewal of certificates used by component changed during upgrades.")
	upgrade_applyCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	upgrade_applyCmd.Flags().Bool("dry-run", false, "Do not change any state, just output what actions would be performed.")
	upgrade_applyCmd.Flags().Bool("etcd-upgrade", true, "Perform the upgrade of etcd.")
	upgrade_applyCmd.Flags().String("feature-gates", "", "A set of key=value pairs that describe feature gates for various features.")
	upgrade_applyCmd.Flags().BoolP("force", "f", false, "Force upgrading although some requirements might not be met. This also implies non-interactive mode.")
	upgrade_applyCmd.Flags().StringSlice("ignore-preflight-errors", []string{}, "A list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.")
	upgrade_applyCmd.Flags().String("kubeconfig", "/etc/kubernetes/admin.conf", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	upgrade_applyCmd.Flags().String("patches", "", "Path to a directory that contains files named \"target[suffix][+patchtype].extension\".")
	upgrade_applyCmd.Flags().Bool("print-config", false, "Specifies whether the configuration file that will be used in the upgrade should be printed or not.")
	upgrade_applyCmd.Flags().BoolP("yes", "y", false, "Perform the upgrade and do not prompt for confirmation (non-interactive mode).")
	upgradeCmd.AddCommand(upgrade_applyCmd)

	carapace.Gen(upgrade_applyCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
		"ignore-preflight-errors": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionChecks().Invoke(c).Filter(c.Parts).ToA()
		}),
		"kubeconfig": carapace.ActionFiles(),
		"patches":    carapace.ActionDirectories(),
	})
}
