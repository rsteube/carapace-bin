package cmd

import (
	"github.com/spf13/cobra"
)

var mod_graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Print a module dependency graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	modCmd.AddCommand(mod_graphCmd)
}
