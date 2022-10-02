package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var profile_enable2faCmd = &cobra.Command{
	Use:   "enable-2fa",
	Short: "enable two-factor authorization",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profile_enable2faCmd).Standalone()

	profileCmd.AddCommand(profile_enable2faCmd)

	carapace.Gen(profile_enable2faCmd).PositionalCompletion(
		carapace.ActionValues("auth-only", "auth-and-writes"),
	)
}
