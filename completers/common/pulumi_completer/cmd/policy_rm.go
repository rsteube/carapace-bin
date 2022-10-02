package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var policy_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes a Policy Pack from a Pulumi organization",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(policy_rmCmd).Standalone()
	policy_rmCmd.PersistentFlags().BoolP("yes", "y", false, "Skip confirmation prompts, and proceed with removal anyway")
	policyCmd.AddCommand(policy_rmCmd)
}
