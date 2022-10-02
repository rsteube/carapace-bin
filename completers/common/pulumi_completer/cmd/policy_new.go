package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var policy_newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new Pulumi Policy Pack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(policy_newCmd).Standalone()
	policy_newCmd.PersistentFlags().String("dir", "", "The location to place the generated Policy Pack; if not specified, the current directory is used")
	policy_newCmd.PersistentFlags().BoolP("force", "f", false, "Forces content to be generated even if it would change existing files")
	policy_newCmd.PersistentFlags().BoolP("generate-only", "g", false, "Generate the Policy Pack only; do not install dependencies")
	policy_newCmd.PersistentFlags().BoolP("offline", "o", false, "Use locally cached templates without making any network requests")
	policyCmd.AddCommand(policy_newCmd)
}
