package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var confirmBeforeCmd = &cobra.Command{
	Use:   "confirm-before",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(confirmBeforeCmd).Standalone()

	confirmBeforeCmd.Flags().StringS("p", "p", "", "prompt")
	confirmBeforeCmd.Flags().StringS("t", "t", "", "target-client")
	rootCmd.AddCommand(confirmBeforeCmd)
}
