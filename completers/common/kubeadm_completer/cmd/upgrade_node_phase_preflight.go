package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubeadm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var upgrade_node_phase_preflightCmd = &cobra.Command{
	Use:   "preflight",
	Short: "Run upgrade node pre-flight checks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_node_phase_preflightCmd).Standalone()
	upgrade_node_phase_preflightCmd.Flags().StringSlice("ignore-preflight-errors", []string{}, "A list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.")
	upgrade_node_phaseCmd.AddCommand(upgrade_node_phase_preflightCmd)

	carapace.Gen(upgrade_node_phase_preflightCmd).FlagCompletion(carapace.ActionMap{
		"ignore-preflight-errors": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionChecks().Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
