package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var label_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List labels in repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(label_listCmd).Standalone()
	label_listCmd.Flags().IntP("page", "p", 1, "Page number")
	label_listCmd.Flags().IntP("per-page", "P", 30, "Number of items to list per page")
	labelCmd.AddCommand(label_listCmd)
}
