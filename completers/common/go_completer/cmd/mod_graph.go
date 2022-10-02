package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var mod_graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "print module requirement graph",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_graphCmd).Standalone()

	modCmd.AddCommand(mod_graphCmd)
}
