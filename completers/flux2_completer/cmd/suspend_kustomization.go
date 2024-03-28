package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var suspend_kustomizationCmd = &cobra.Command{
	Use:   "kustomization",
	Short: "Suspend reconciliation of Kustomization",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suspend_kustomizationCmd).Standalone()
	suspendCmd.AddCommand(suspend_kustomizationCmd)
}
