package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var join_phaseCmd = &cobra.Command{
	Use:   "phase",
	Short: "Use this command to invoke single phase of the join workflow",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(join_phaseCmd).Standalone()
	joinCmd.AddCommand(join_phaseCmd)
}
