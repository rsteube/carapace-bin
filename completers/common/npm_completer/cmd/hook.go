package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Manage registry hooks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hookCmd).Standalone()
	hookCmd.PersistentFlags().String("otp", "", "one-time password")

	rootCmd.AddCommand(hookCmd)
}
