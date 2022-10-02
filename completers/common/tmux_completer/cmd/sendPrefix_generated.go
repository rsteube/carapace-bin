package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var sendPrefixCmd = &cobra.Command{
	Use:   "send-prefix",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sendPrefixCmd).Standalone()

	sendPrefixCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(sendPrefixCmd)
}
