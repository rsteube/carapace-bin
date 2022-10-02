package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubeadm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Performs a best effort revert of changes made to this host by 'kubeadm init' or 'kubeadm join'",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resetCmd).Standalone()
	resetCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path to the directory where the certificates are stored. If specified, clean this directory.")
	resetCmd.Flags().String("cri-socket", "", "Path to the CRI socket to connect. If empty kubeadm will try to auto-detect this value; use this option only if you have more than one CRI installed or if you have non-standard CRI socket.")
	resetCmd.Flags().BoolP("force", "f", false, "Reset the node without prompting for confirmation.")
	resetCmd.Flags().StringSlice("ignore-preflight-errors", []string{}, "A list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.")
	resetCmd.Flags().String("kubeconfig", "/etc/kubernetes/admin.conf", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	resetCmd.Flags().StringSlice("skip-phases", []string{}, "List of phases to be skipped")
	rootCmd.AddCommand(resetCmd)

	carapace.Gen(resetCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir":   carapace.ActionDirectories(),
		"cri-socket": carapace.ActionFiles(),
		"ignore-preflight-errors": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
				return action.ActionChecks().Invoke(c).Filter(c.Parts).ToA()
			})
		}),
		"kubeconfig": carapace.ActionFiles(),
		"skip-phases": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionPhases().Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
