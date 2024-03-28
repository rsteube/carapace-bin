package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var suspend_sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Suspend sources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suspend_sourceCmd).Standalone()
	suspendCmd.AddCommand(suspend_sourceCmd)
}
