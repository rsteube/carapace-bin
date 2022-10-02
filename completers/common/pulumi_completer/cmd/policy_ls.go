package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var policy_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all Policy Packs for a Pulumi organization",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(policy_lsCmd).Standalone()
	policy_lsCmd.PersistentFlags().BoolP("json", "j", false, "Emit output as JSON")
	policyCmd.AddCommand(policy_lsCmd)
}
