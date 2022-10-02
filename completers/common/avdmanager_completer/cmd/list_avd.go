package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var list_avdCmd = &cobra.Command{
	Use:   "avd",
	Short: "Lists existing Android Virtual Devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(list_avdCmd).Standalone()

	list_avdCmd.Flags().BoolP("compact", "c", false, "Compact output (suitable for scripts)")
	list_avdCmd.Flags().BoolP("null", "0", false, "Terminates lines with \\0 instead of \\n (e.g. for xargs -0).")
	listCmd.AddCommand(list_avdCmd)
}
