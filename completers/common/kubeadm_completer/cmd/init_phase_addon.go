package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_addonCmd = &cobra.Command{
	Use:   "addon",
	Short: "Install required addons for passing conformance tests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_addonCmd).Standalone()
	init_phaseCmd.AddCommand(init_phase_addonCmd)
}
