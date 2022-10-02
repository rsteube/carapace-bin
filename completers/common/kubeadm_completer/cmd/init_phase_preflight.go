package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubeadm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var init_phase_preflightCmd = &cobra.Command{
	Use:   "preflight",
	Short: "Run pre-flight checks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_preflightCmd).Standalone()
	init_phase_preflightCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_preflightCmd.Flags().StringSlice("ignore-preflight-errors", []string{}, "A list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.")
	init_phaseCmd.AddCommand(init_phase_preflightCmd)

	carapace.Gen(init_phase_preflightCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
		"ignore-preflight-errors": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionChecks().Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
