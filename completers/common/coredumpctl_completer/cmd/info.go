package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show detailed information about one or more coredumps",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	rootCmd.AddCommand(infoCmd)
}
