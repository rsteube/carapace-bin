package cmd

import (
	"github.com/spf13/cobra"
)

var list_expiredCmd = &cobra.Command{
	Use:   "expired",
	Short: "List all posts already expired",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	listCmd.AddCommand(list_expiredCmd)
}
