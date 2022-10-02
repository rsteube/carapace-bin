package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var list_deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Lists existing devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(list_deviceCmd).Standalone()

	list_deviceCmd.Flags().BoolP("compact", "c", false, "Compact output (suitable for scripts)")
	list_deviceCmd.Flags().BoolP("null", "0", false, "Terminates lines with \\0 instead of \\n (e.g. for xargs -0).")
	listCmd.AddCommand(list_deviceCmd)
}
