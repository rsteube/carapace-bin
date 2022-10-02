package cmd

import (
	"github.com/spf13/cobra"
)

var write_treeCmd = &cobra.Command{
	Use:   "write-tree",
	Short: "Create a tree object from the current index",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	write_treeCmd.Flags().Bool("missing-ok", false, "allow missing objects")
	write_treeCmd.Flags().String("prefix", "", "write tree object for a subdirectory <prefix>")
	rootCmd.AddCommand(write_treeCmd)
}
