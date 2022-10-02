package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var downgradeCmd = &cobra.Command{
	Use:   "downgrade",
	Short: "Downgrade Flutter to the last active version for the current channel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downgradeCmd).Standalone()

	downgradeCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	rootCmd.AddCommand(downgradeCmd)
}
