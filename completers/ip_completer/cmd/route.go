package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var routeCmd = &cobra.Command{
	Use:   "route",
	Short: "routing table entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(routeCmd).Standalone()

	rootCmd.AddCommand(routeCmd)
}
