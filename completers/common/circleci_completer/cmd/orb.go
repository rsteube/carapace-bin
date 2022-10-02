package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var orbCmd = &cobra.Command{
	Use:   "orb",
	Short: "Operate on orbs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orbCmd).Standalone()
	rootCmd.AddCommand(orbCmd)
}
