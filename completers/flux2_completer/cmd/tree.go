package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Print the resources reconciled by Flux",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(treeCmd).Standalone()
	rootCmd.AddCommand(treeCmd)
}
