package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubeadm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var upgrade_planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Check which versions are available to upgrade to and validate whether your current cluster is upgradeable. To skip the internet check, pass in the optional [version] parameter",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_planCmd).Standalone()
	upgrade_planCmd.Flags().Bool("allow-experimental-upgrades", false, "Show unstable versions of Kubernetes as an upgrade alternative and allow upgrading to an alpha/beta/release candidate versions of Kubernetes.")
	upgrade_planCmd.Flags().Bool("allow-release-candidate-upgrades", false, "Show release candidate versions of Kubernetes as an upgrade alternative and allow upgrading to a release candidate versions of Kubernetes.")
	upgrade_planCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	upgrade_planCmd.Flags().String("feature-gates", "", "A set of key=value pairs that describe feature gates for various features.")
	upgrade_planCmd.Flags().StringSlice("ignore-preflight-errors", []string{}, "A list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.")
	upgrade_planCmd.Flags().String("kubeconfig", "/etc/kubernetes/admin.conf", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	upgrade_planCmd.Flags().Bool("print-config", false, "Specifies whether the configuration file that will be used in the upgrade should be printed or not.")
	upgradeCmd.AddCommand(upgrade_planCmd)

	carapace.Gen(upgrade_planCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
		"ignore-preflight-errors": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionChecks().Invoke(c).Filter(c.Parts).ToA()
		}),
		"kubeconfig": carapace.ActionFiles(),
	})
}
