package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var ownerCmd = &cobra.Command{
	Use:   "owner",
	Short: "Manage package owners",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ownerCmd).Standalone()
	ownerCmd.PersistentFlags().String("otp", "", "one-time password")

	rootCmd.AddCommand(ownerCmd)
}
