package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var keygenCmd = &cobra.Command{
	Use:   "keygen",
	Short: "generate adb public/private key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keygenCmd).Standalone()

	rootCmd.AddCommand(keygenCmd)

	carapace.Gen(keygenCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
