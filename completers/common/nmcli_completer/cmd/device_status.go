package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var device_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status for all devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_statusCmd).Standalone()

	deviceCmd.AddCommand(device_statusCmd)
}
