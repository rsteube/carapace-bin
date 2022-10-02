package cmd

import (
	"github.com/spf13/cobra"
)

var list_draftsCmd = &cobra.Command{
	Use:   "drafts",
	Short: "List all drafts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	listCmd.AddCommand(list_draftsCmd)
}
