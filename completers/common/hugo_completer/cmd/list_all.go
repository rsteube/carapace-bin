package cmd

import (
	"github.com/spf13/cobra"
)

var list_allCmd = &cobra.Command{
	Use:   "all",
	Short: "List all posts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	listCmd.AddCommand(list_allCmd)
}
