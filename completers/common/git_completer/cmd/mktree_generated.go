package cmd

import (
	"github.com/spf13/cobra"
)

var mktreeCmd = &cobra.Command{
	Use:   "mktree",
	Short: "Build a tree-object from ls-tree formatted text",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	mktreeCmd.Flags().Bool("batch", false, "allow creation of more than one tree")
	mktreeCmd.Flags().Bool("missing", false, "allow missing objects")
	mktreeCmd.Flags().BoolS("z", "z", false, "input is NUL terminated")
	rootCmd.AddCommand(mktreeCmd)
}
