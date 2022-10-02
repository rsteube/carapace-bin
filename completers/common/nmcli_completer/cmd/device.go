package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "devices managed by NetworkManager",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deviceCmd).Standalone()

	rootCmd.AddCommand(deviceCmd)
}
