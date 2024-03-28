package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var tunnelCmd = &cobra.Command{
	Use:   "tunnel",
	Short: "tunnel over IP",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tunnelCmd).Standalone()

	rootCmd.AddCommand(tunnelCmd)
}
