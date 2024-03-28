package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var create_receiverCmd = &cobra.Command{
	Use:   "receiver",
	Short: "Create or update a Receiver resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_receiverCmd).Standalone()
	create_receiverCmd.Flags().StringSlice("event", []string{}, "also accepts comma-separated values")
	create_receiverCmd.Flags().StringSlice("resource", []string{}, "also accepts comma-separated values")
	create_receiverCmd.Flags().String("secret-ref", "", "")
	create_receiverCmd.Flags().String("type", "", "")
	createCmd.AddCommand(create_receiverCmd)
}
