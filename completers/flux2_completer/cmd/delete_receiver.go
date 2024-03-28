package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var delete_receiverCmd = &cobra.Command{
	Use:   "receiver",
	Short: "Delete a Receiver resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delete_receiverCmd).Standalone()
	deleteCmd.AddCommand(delete_receiverCmd)
}
