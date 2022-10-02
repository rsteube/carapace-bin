package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var suspendClientCmd = &cobra.Command{
	Use:   "suspend-client",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suspendClientCmd).Standalone()

	suspendClientCmd.Flags().StringS("t", "t", "", "target-client")
	rootCmd.AddCommand(suspendClientCmd)
}
