package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phaseCmd = &cobra.Command{
	Use:   "phase",
	Short: "Use this command to invoke single phase of the init workflow",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phaseCmd).Standalone()
	initCmd.AddCommand(init_phaseCmd)
}
