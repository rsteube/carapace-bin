package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var suspend_source_helmCmd = &cobra.Command{
	Use:   "helm",
	Short: "Suspend reconciliation of a HelmRepository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suspend_source_helmCmd).Standalone()
	suspend_sourceCmd.AddCommand(suspend_source_helmCmd)
}
