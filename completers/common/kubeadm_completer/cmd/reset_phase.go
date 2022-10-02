package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reset_phaseCmd = &cobra.Command{
	Use:   "phase",
	Short: "Use this command to invoke single phase of the reset workflow",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reset_phaseCmd).Standalone()
	resetCmd.AddCommand(reset_phaseCmd)
}
