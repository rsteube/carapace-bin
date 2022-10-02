package cmd

import (
	"github.com/spf13/cobra"
)

var commit_treeCmd = &cobra.Command{
	Use:   "commit-tree",
	Short: "Create a new commit object",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	commit_treeCmd.Flags().StringS("F", "F", "", "read commit log message from file")
	commit_treeCmd.Flags().StringP("gpg-sign", "S", "", "GPG sign commit")
	commit_treeCmd.Flags().StringS("m", "m", "", "commit message")
	commit_treeCmd.Flags().StringS("p", "p", "", "id of a parent commit object")
	rootCmd.AddCommand(commit_treeCmd)
}
