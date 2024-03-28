package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var suspend_source_chartCmd = &cobra.Command{
	Use:   "chart",
	Short: "Suspend reconciliation of a HelmChart",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suspend_source_chartCmd).Standalone()
	suspend_sourceCmd.AddCommand(suspend_source_chartCmd)
}
