package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var policy_group_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all Policy Groups for a Pulumi organization",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(policy_group_lsCmd).Standalone()
	policy_group_lsCmd.PersistentFlags().BoolP("json", "j", false, "Emit output as JSON")
	policy_groupCmd.AddCommand(policy_group_lsCmd)
}
