package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var usbCmd = &cobra.Command{
	Use:   "usb",
	Short: "restart adbd listening on USB",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(usbCmd).Standalone()

	rootCmd.AddCommand(usbCmd)
}
