package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_certsCmd = &cobra.Command{
	Use:   "certs",
	Short: "Certificate generation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_certsCmd).Standalone()
	init_phaseCmd.AddCommand(init_phase_certsCmd)
}
