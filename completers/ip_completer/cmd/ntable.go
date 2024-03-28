package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var ntableCmd = &cobra.Command{
	Use:   "ntable",
	Short: "manage the neighbor cache's operation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ntableCmd).Standalone()

	rootCmd.AddCommand(ntableCmd)
}
