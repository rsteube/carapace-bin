package cmd

import (
	"github.com/spf13/cobra"
)

var node_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a node from a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	nodeCmd.AddCommand(node_deleteCmd)
}
