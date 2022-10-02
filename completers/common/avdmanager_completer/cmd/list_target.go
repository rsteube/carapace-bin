package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var list_targetCmd = &cobra.Command{
	Use:   "target",
	Short: "Lists existing targets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(list_targetCmd).Standalone()

	list_targetCmd.Flags().BoolP("compact", "c", false, "Compact output (suitable for scripts)")
	list_targetCmd.Flags().BoolP("null", "0", false, "Terminates lines with \\0 instead of \\n (e.g. for xargs -0).")
	listCmd.AddCommand(list_targetCmd)
}
