package cmd

import (
	"github.com/spf13/cobra"
)

var node_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	node_startCmd.Flags().Bool("delete-on-failure", false, "If set, delete the current cluster if start fails and try again. Defaults to false.")
	nodeCmd.AddCommand(node_startCmd)
}
