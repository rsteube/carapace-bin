package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var sendKeysCmd = &cobra.Command{
	Use:   "send-keys",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sendKeysCmd).Standalone()

	sendKeysCmd.Flags().BoolS("F", "F", false, "TODO description")
	sendKeysCmd.Flags().BoolS("H", "H", false, "TODO description")
	sendKeysCmd.Flags().BoolS("M", "M", false, "TODO description")
	sendKeysCmd.Flags().StringS("N", "N", "", "repeat-count")
	sendKeysCmd.Flags().BoolS("R", "R", false, "TODO description")
	sendKeysCmd.Flags().BoolS("X", "X", false, "TODO description")
	sendKeysCmd.Flags().BoolS("l", "l", false, "TODO description")
	sendKeysCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(sendKeysCmd)
}
