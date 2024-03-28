package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var export_receiverCmd = &cobra.Command{
	Use:   "receiver",
	Short: "Export Receiver resources in YAML format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(export_receiverCmd).Standalone()
	exportCmd.AddCommand(export_receiverCmd)
}
