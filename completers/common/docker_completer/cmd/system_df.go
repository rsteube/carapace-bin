package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var system_dfCmd = &cobra.Command{
	Use:   "df",
	Short: "Show docker disk usage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_dfCmd).Standalone()

	system_dfCmd.Flags().String("format", "", "Pretty-print images using a Go template")
	system_dfCmd.Flags().BoolP("verbose", "v", false, "Show detailed information on space usage")
	systemCmd.AddCommand(system_dfCmd)
}
