package cmd

import (
	"github.com/spf13/cobra"
)

var node_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List nodes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	nodeCmd.AddCommand(node_listCmd)
}
