package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var showMessagesCmd = &cobra.Command{
	Use:   "show-messages",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showMessagesCmd).Standalone()

	showMessagesCmd.Flags().BoolS("J", "J", false, "TODO description")
	showMessagesCmd.Flags().BoolS("T", "T", false, "TODO description")
	showMessagesCmd.Flags().StringS("t", "t", "", "target-client")
	rootCmd.AddCommand(showMessagesCmd)
}
