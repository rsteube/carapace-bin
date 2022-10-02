package cmd

import (
	"github.com/spf13/cobra"
)

var merge_treeCmd = &cobra.Command{
	Use:   "merge-tree",
	Short: "Show three-way merge without touching index",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	rootCmd.AddCommand(merge_treeCmd)
}
