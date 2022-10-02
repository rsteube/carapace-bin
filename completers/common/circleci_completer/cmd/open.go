package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open the current project in the browser.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(openCmd).Standalone()
	rootCmd.AddCommand(openCmd)
}
