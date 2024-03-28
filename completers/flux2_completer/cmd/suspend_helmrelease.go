package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var suspend_helmreleaseCmd = &cobra.Command{
	Use:   "helmrelease",
	Short: "Suspend reconciliation of HelmRelease",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suspend_helmreleaseCmd).Standalone()
	suspendCmd.AddCommand(suspend_helmreleaseCmd)
}
