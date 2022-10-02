package cmd

import (
	"github.com/spf13/cobra"
)

var commit_graphCmd = &cobra.Command{
	Use:   "commit-graph",
	Short: "Write and verify Git commit-graph files",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	commit_graphCmd.Flags().String("object-dir", "", "The object directory to store the graph")
	rootCmd.AddCommand(commit_graphCmd)
}
