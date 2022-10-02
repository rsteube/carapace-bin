package cmd

import (
	"github.com/spf13/cobra"
)

var list_futureCmd = &cobra.Command{
	Use:   "future",
	Short: "List all posts dated in the future",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	listCmd.AddCommand(list_futureCmd)
}
