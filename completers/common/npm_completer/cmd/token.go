package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Manage your authentication tokens",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tokenCmd).Standalone()

	tokenCmd.PersistentFlags().String("otp", "", "one-time password")
	rootCmd.AddCommand(tokenCmd)
}
