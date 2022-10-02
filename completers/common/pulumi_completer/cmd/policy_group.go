package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var policy_groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage policy groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(policy_groupCmd).Standalone()
	policyCmd.AddCommand(policy_groupCmd)
}
