package cmd

import (
	"github.com/spf13/cobra"
)

var node_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a node to the given cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	node_addCmd.Flags().Bool("control-plane", false, "If true, the node added will also be a control plane in addition to a worker.")
	node_addCmd.Flags().Bool("delete-on-failure", false, "If set, delete the current cluster if start fails and try again. Defaults to false.")
	node_addCmd.Flags().Bool("worker", true, "If true, the added node will be marked for work. Defaults to true.")
	nodeCmd.AddCommand(node_addCmd)
}
